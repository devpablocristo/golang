package org.example;

import org.eclipse.jgit.api.Git;
import org.eclipse.jgit.api.errors.GitAPIException;
import org.eclipse.jgit.lib.Repository;
import org.eclipse.jgit.revwalk.RevCommit;

import java.io.File;
import java.io.IOException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Iterator;

public class GitUtils {

    public static String getFileAuthor(Repository repo, String relativePath) throws IOException, GitAPIException {
        try (Git git = new Git(repo)) {
            //String relativePath = getRelativePath(repo, file);

            Iterable<RevCommit> commits = git.log().addPath(relativePath).call();
            Iterator<RevCommit> commitIterator = commits.iterator();

            if (commitIterator.hasNext()) {
                RevCommit firstCommit = commitIterator.next();
                return firstCommit.getAuthorIdent().getEmailAddress();
            } else {
                System.out.println("No se encontraron commits para el archivo: " + relativePath);
                return "Autor desconocido";
            }
        }
    }

    public static String getCommitID(Repository repo, String relativePath) throws IOException, GitAPIException {
        try (Git git = new Git(repo)) {
            //String relativePath = getRelativePath(repo, file);

            Iterable<RevCommit> commits = git.log().addPath(relativePath).call();
            Iterator<RevCommit> commitIterator = commits.iterator();

            if (commitIterator.hasNext()) {
                RevCommit firstCommit = commitIterator.next();
                return firstCommit.getName();
            }

            System.out.println("No se encontraron commits para el archivo: " + relativePath);
            return "id not found";
        }
    }

    public static String getRelativePath(Repository repo, String file){
        File repoDirectory = repo.getDirectory().getParentFile();
        Path relativePath = repoDirectory.toPath().relativize(Paths.get(file));

        return relativePath.toString();
    }
}

