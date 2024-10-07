package org.example;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.eclipse.jgit.api.errors.GitAPIException;
import org.eclipse.jgit.lib.Repository;
import org.eclipse.jgit.storage.file.FileRepositoryBuilder;
import com.github.javaparser.StaticJavaParser;
import com.github.javaparser.ast.CompilationUnit;
import com.github.javaparser.ast.body.ClassOrInterfaceDeclaration;
import com.github.javaparser.ast.body.MethodDeclaration;
import com.github.javaparser.ast.visitor.VoidVisitorAdapter;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.*;

public class MetricsAnalyzer {
    private final Repository repo;
    private final Map<String, Metric> metricsMap = new HashMap<>();
    private static final List<Skill> skills = Arrays.asList(
            Skill.METHOD_LENGTH,
            Skill.NAMING_CONVENTIONS,
            Skill.CYCLOMATIC_COMPLEXITY,
            Skill.METHOD_PARAMETERS);

    public MetricsAnalyzer(Repository repo) {
        this.repo = repo;
    }

    public List<Metric> getMetrics() {
        return new ArrayList<>(metricsMap.values());
    }

    public enum Skill {
        METHOD_LENGTH("method_length", "Evaluates the length of methods in the code to ensure they do not exceed a set limit, improving code readability and maintainability."),
        NAMING_CONVENTIONS("naming_conventions", "Checks that class and method names follow standard naming conventions, promoting consistency in the code."),
        CYCLOMATIC_COMPLEXITY("cyclomatic_complexity", "Checks that class and method names follow standard naming conventions, promoting consistency in the code."),
        METHOD_PARAMETERS("method_parameters", "Checks the number of parameters in methods to ensure methods are not overly complex.");

        private final String id;
        private final String description;

        Skill(String id, String description) {
            this.id = id;
            this.description = description;
        }

        public String getId() {
            return id;
        }
    }

    public static class Evidence {
        public String commitID;
        public String file;
        public int line;
    }

    public static class Metric {
        public String skillID;
        public String gitAuthor;
        public String score;
        public List<Evidence> evidence;
    }

    private class NamingConventionAnalyzer extends VoidVisitorAdapter<Void> {
        private final String file;

        public NamingConventionAnalyzer(String file) {
            this.file = file;
        }

        @Override
        public void visit(ClassOrInterfaceDeclaration cid, Void arg) {
            super.visit(cid, arg);
            String className = cid.getNameAsString();
            int lineNumber = cid.getBegin().get().line;

            if (!Character.isUpperCase(className.charAt(0))) {
                addMetric(Skill.NAMING_CONVENTIONS.id, lineNumber, file);
            } else {
                addMetric(Skill.NAMING_CONVENTIONS.id, -1, file);
            }

            cid.getMethods().forEach(this::analyzeMethodNaming);
        }

        private void analyzeMethodNaming(MethodDeclaration md) {
            String methodName = md.getNameAsString();
            int lineNumber = md.getBegin().get().line;
            if (!Character.isLowerCase(methodName.charAt(0))) {
                addMetric(Skill.NAMING_CONVENTIONS.id, lineNumber, file);
            } else {
                addMetric(Skill.NAMING_CONVENTIONS.id, -1, file);
            }
        }
    }

    private class MethodLengthAnalyzer extends VoidVisitorAdapter<Void> {
        private static final int MAX_METHOD_LENGTH = 30;
        private final String file;

        public MethodLengthAnalyzer(String file) {
            this.file = file;
        }

        @Override
        public void visit(MethodDeclaration md, Void arg) {
            if (file.endsWith("Test.java")) {
                return;
            }

            super.visit(md, arg);
            int beginLine = md.getBegin().get().line;
            int methodLength = md.getEnd().get().line - beginLine;

            if (methodLength > MAX_METHOD_LENGTH) {
                addMetric(Skill.METHOD_LENGTH.id, beginLine, file);
            }
        }
    }

    private class CyclomaticComplexityAnalyzer extends VoidVisitorAdapter<Void> {
        private final String file;

        public CyclomaticComplexityAnalyzer(String file) {
            this.file = file;
        }

        @Override
        public void visit(MethodDeclaration md, Void arg) {
            super.visit(md, arg);
            int beginLine = md.getBegin().get().line;

            int cyclomaticComplexity = calculateCyclomaticComplexity(md);

            int THRESHOLD = 10;
            if (cyclomaticComplexity > THRESHOLD) {
                addMetric(Skill.CYCLOMATIC_COMPLEXITY.id, beginLine, file);
            } else {
                addMetric(Skill.CYCLOMATIC_COMPLEXITY.id, -1, file); // No hay evidencia, puntaje queda en 1
            }
        }

        private int calculateCyclomaticComplexity(MethodDeclaration md) {
            int complexity = 1;
            String methodBody = md.toString();

            complexity += countOccurrences(methodBody, "if");
            complexity += countOccurrences(methodBody, "for");
            complexity += countOccurrences(methodBody, "while");
            complexity += countOccurrences(methodBody, "case");
            complexity += countOccurrences(methodBody, "catch");
            complexity += countOccurrences(methodBody, "&&");
            complexity += countOccurrences(methodBody, "||");

            return complexity;
        }

        private int countOccurrences(String body, String keyword) {
            int count = 0;
            int index = 0;
            while ((index = body.indexOf(keyword, index)) != -1) {
                count++;
                index += keyword.length();
            }
            return count;
        }
    }

    private class MethodParametersAnalyzer extends VoidVisitorAdapter<Void> {
        private static final int MAX_PARAMETERS = 4;
        private final String file;

        public MethodParametersAnalyzer(String file) {
            this.file = file;
        }

        @Override
        public void visit(MethodDeclaration md, Void arg) {
            super.visit(md, arg);

            if (file.endsWith("Test.java")) {
                return;
            }

            int parameterCount = md.getParameters().size();
            int lineNumber = md.getBegin().get().line;

            if (parameterCount > MAX_PARAMETERS) {
                addMetric(Skill.METHOD_PARAMETERS.id, lineNumber, file);
            } else {
                addMetric(Skill.METHOD_PARAMETERS.id, -1, file);
            }
        }
    }

    private void addMetric(String metricId, int lineNumber, String filePath) {
        try {
            String author = GitUtils.getFileAuthor(repo, filePath);
            String commitId = GitUtils.getCommitID(repo, filePath);

            String key = metricId + "|" + author;

            if (lineNumber == -1) {
                if (!metricsMap.containsKey(key)) {
                    Metric metric = new Metric();
                    metric.skillID = metricId;
                    metric.gitAuthor = author;
                    metric.score = "1";
                    metric.evidence = new ArrayList<>();

                    metricsMap.put(key, metric);
                }
                return;
            }

            Evidence evidence = new Evidence();
            evidence.commitID = commitId;
            evidence.file = filePath;
            evidence.line = lineNumber;

            if (metricsMap.containsKey(key)) {
                Metric existingMetric = metricsMap.get(key);
                existingMetric.evidence.add(evidence);
                existingMetric.score = "0";
            } else {
                Metric metric = new Metric();
                metric.skillID = metricId;
                metric.gitAuthor = author;
                metric.score = "0";
                metric.evidence = new ArrayList<>();
                metric.evidence.add(evidence);

                metricsMap.put(key, metric);
            }

        } catch (IOException | GitAPIException e) {
            System.out.println("Error obteniendo la información de Git: " + e.getMessage());
        }
    }

    public void analyzeFile(File file) throws IOException {
        try (FileInputStream in = new FileInputStream(file)){
            CompilationUnit cu = StaticJavaParser.parse(in);

            String path = GitUtils.getRelativePath(repo, file.getPath());

            for (Skill skill : skills) {
                switch (skill.id) {
                    case "method_length":
                        MethodLengthAnalyzer methodLengthAnalyzer = new MethodLengthAnalyzer(path);
                        methodLengthAnalyzer.visit(cu, null);
                        break;
                    case "naming_conventions":
                        NamingConventionAnalyzer namingAnalyzer = new NamingConventionAnalyzer(path);
                        namingAnalyzer.visit(cu, null);
                        break;
                    case "cyclomatic_complexity":
                        CyclomaticComplexityAnalyzer complexityAnalyzer = new CyclomaticComplexityAnalyzer(path);
                        complexityAnalyzer.visit(cu, null);
                        break;
                    case "method_parameters":
                        MethodParametersAnalyzer parametersAnalyzer = new MethodParametersAnalyzer(path);
                        parametersAnalyzer.visit(cu, null);
                        break;
                    default:
                        System.out.printf("skill not implemented: %s%n", skill.id);
                }
            }
        } catch (IOException e) {
            System.out.println("Error al analizar el archivo: " + file.getName());
            throw e;
        }
    }

    public static void main(String[] args) {
        if (args.length == 0) {
            System.out.println("Por favor, proporciona la ruta al repositorio.");
            return;
        }

        String repoPath = args[0];

        try {
            Repository repo = new FileRepositoryBuilder()
                    .setGitDir(new File(repoPath))
                    .readEnvironment()
                    .findGitDir()
                    .build();

            MetricsAnalyzer analyzer = new MetricsAnalyzer(repo);

            Path rootPath = Paths.get(repoPath).getParent();
            Files.walk(rootPath)
                    .filter(Files::isRegularFile)
                    .filter(path -> path.toString().endsWith(".java"))
                    .forEach(path -> {
                        try {
                            analyzer.analyzeFile(path.toFile());
                        } catch (IOException e) {
                            System.out.println("Error al analizar archivo: " + path + " - " + e.getMessage());
                        }
                    });

            ObjectMapper objectMapper = new ObjectMapper();
            String jsonOutput = objectMapper.writerWithDefaultPrettyPrinter().writeValueAsString(analyzer.getMetrics());
            System.out.println(jsonOutput);

        } catch (IOException e) {
            System.out.println("Error al abrir el repositorio: " + e.getMessage());
        }
    }
}
