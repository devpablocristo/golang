package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	handler "compile-daemon/infra/handler"
)

func main() {
	// Configuración de AWS DynamoDB
	config := &aws.Config{
		Endpoint: aws.String("http://dynamodb-local:8765"), // Nombre del servicio y puerto de Docker Compose
		Region:   aws.String("local"),                      // Usa 'local' para DynamoDB local
		Credentials: credentials.NewStaticCredentials(
			"fakeAccessKey",
			"fakeSecretKey",
			"",
		),
	}

	// Crear una nueva sesión
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}

	// Crear un cliente DynamoDB
	svc := dynamodb.New(sess)

	// Ejemplo: Listar tablas
	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Panicf("error: %v", err)
	}

	fmt.Println("Tablas en DynamoDB:")
	for _, table := range result.TableNames {
		fmt.Println(*table)
	}

	http.HandleFunc("/", handler.HomePage)
	http.HandleFunc("/users", handler.UserPage)
	// puerto del contenedor
	log.Fatal(http.ListenAndServe(":8888", nil))
}

// package main

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"time"

// 	"github.com/aws/aws-sdk-go/aws"
// 	"github.com/aws/aws-sdk-go/aws/awserr"
// 	"github.com/aws/aws-sdk-go/aws/credentials"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/dynamodb"
// 	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
// 	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
// 	"github.com/google/uuid"
// 	"github.com/joho/godotenv"

// 	app "github.com/teamcubation/TeamcuBot-Back/cmd/server/app"
// 	ctypes "github.com/teamcubation/TeamcuBot-Back/internal/platform/custom-types"
// )

// type Data struct {
// 	ID          string
// 	Temperature int32
// 	Llm         string
// 	Rookie      string
// 	Prompt      string
// 	Context     string
// 	Process     string
// 	DateFrom    *time.Time
// 	DateTo      *time.Time
// }

// type Movie struct {
// 	Year   int
// 	Title  string
// 	Plot   string
// 	Rating float64
// }

// func main() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	if err := godotenv.Load(); err != nil {
// 		ctypes.HandleFatalError("error loading .env file", err)
// 	}

// 	// Create a DynamoDB client
// 	svc := dynamodbConfig()

// 	// DynamoDBCreateItem
// 	// DynamoDBCreateTable
// 	// DynamoDBDeleteItem
// 	// DynamoDBListTables
// 	// DynamoDBLoadItems
// 	// DynamoDBReadItem
// 	// DynamoDBScanItems
// 	// DynamoDBUpdateItem

// 	//createTableData(svc)
// 	//createTableMovies(svc)

// 	//createDataItem(svc)
// 	//createMovieItem(svc)

// 	//readDataItem(svc)
// 	//readMovieItem(svc)

// 	//scanDataItem(svc)
// 	//scanMovieItem(svc)

// 	//deleteDataItem(svc)
// 	//deleteMovieItem(svc)

// 	//updateDataItem(svc)
// 	//updateMovieItem(svc)

// 	loadDataItemsJson(svc)
// 	//loadMovieItemsJson(svc)

// 	//listTables(svc)

// 	app.Build(ctx)
// }

// func createTableData(svc *dynamodb.DynamoDB) {
// 	tableName := "Data2"

// 	input := &dynamodb.CreateTableInput{
// 		AttributeDefinitions: []*dynamodb.AttributeDefinition{
// 			{
// 				AttributeName: aws.String("ID"),
// 				AttributeType: aws.String("S"),
// 			},
// 		},
// 		KeySchema: []*dynamodb.KeySchemaElement{
// 			{
// 				AttributeName: aws.String("ID"),
// 				KeyType:       aws.String("HASH"),
// 			},
// 		},
// 		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
// 			ReadCapacityUnits:  aws.Int64(10),
// 			WriteCapacityUnits: aws.Int64(10),
// 		},
// 		TableName: aws.String(tableName),
// 	}

// 	_, err := svc.CreateTable(input)
// 	if err != nil {
// 		log.Fatalf("Got error calling CreateTable: %s", err)
// 	}

// 	fmt.Println("Created the table", tableName)
// }

// func createTableMovies(svc *dynamodb.DynamoDB) {
// 	tableName := "Movies"

// 	input := &dynamodb.CreateTableInput{
// 		AttributeDefinitions: []*dynamodb.AttributeDefinition{
// 			{
// 				AttributeName: aws.String("Year"),
// 				AttributeType: aws.String("N"),
// 			},
// 			{
// 				AttributeName: aws.String("Title"),
// 				AttributeType: aws.String("S"),
// 			},
// 		},
// 		KeySchema: []*dynamodb.KeySchemaElement{
// 			{
// 				AttributeName: aws.String("Year"),
// 				KeyType:       aws.String("HASH"),
// 			},
// 			{
// 				AttributeName: aws.String("Title"),
// 				KeyType:       aws.String("RANGE"),
// 			},
// 		},
// 		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
// 			ReadCapacityUnits:  aws.Int64(10),
// 			WriteCapacityUnits: aws.Int64(10),
// 		},
// 		TableName: aws.String(tableName),
// 	}

// 	_, err := svc.CreateTable(input)
// 	if err != nil {
// 		log.Fatalf("Got error calling CreateTable: %s", err)
// 	}

// 	fmt.Println("Created the table", tableName)
// }

// func createDataItem(svc *dynamodb.DynamoDB) {
// 	newUUID := uuid.New()

// 	// Convertir el UUID a su representación de cadena (string)
// 	uuidString := newUUID.String()

// 	// DynamoDB solo autogenera valores numericos (autoincremento)

// 	dataItem := Data{
// 		ID:          uuidString,
// 		Temperature: 25,
// 		Llm:         "Sample Llm",
// 		Rookie:      "Sample Rookie",
// 		Prompt:      "Sample Prompt",
// 		Context:     "Sample Context",
// 		Process:     "Sample Process",
// 	}

// 	av, err := dynamodbattribute.MarshalMap(dataItem)
// 	if err != nil {
// 		log.Fatalf("Got error marshalling data item: %s", err)
// 	}

// 	tableName := "Data2"

// 	input := &dynamodb.PutItemInput{
// 		Item:      av,
// 		TableName: aws.String(tableName),
// 	}

// 	_, err = svc.PutItem(input)
// 	if err != nil {
// 		log.Fatalf("Got error calling PutItem: %s", err)
// 	}

// 	fmt.Println("Successfully added item with ID '" + dataItem.ID + "' to table " + tableName)
// }

// func createMovieItem(svc *dynamodb.DynamoDB) {
// 	Movie := Movie{
// 		Year:   2015,
// 		Title:  "The Big New Movie",
// 		Plot:   "Nothing happens at all.",
// 		Rating: 0.0,
// 	}

// 	av, err := dynamodbattribute.MarshalMap(Movie)
// 	if err != nil {
// 		log.Fatalf("Got error marshalling new movie Movie: %s", err)
// 	}

// 	tableName := "Movies"

// 	input := &dynamodb.PutItemInput{
// 		Item:      av,
// 		TableName: aws.String(tableName),
// 	}

// 	_, err = svc.PutItem(input)
// 	if err != nil {
// 		log.Fatalf("Got error calling PutMovie: %s", err)
// 	}

// 	year := strconv.Itoa(Movie.Year)

// 	fmt.Println("Successfully added '" + Movie.Title + "' (" + year + ") to table " + tableName)
// }

// func readDataItem(svc *dynamodb.DynamoDB) {
// 	tableName := "Data2"
// 	itemID := "12345"

// 	result, err := svc.GetItem(&dynamodb.GetItemInput{
// 		TableName: aws.String(tableName),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"ID": {
// 				S: aws.String(itemID),
// 			},
// 		},
// 	})
// 	if err != nil {
// 		log.Fatalf("Got error calling GetItem: %s", err)
// 	}

// 	if result.Item == nil {
// 		log.Fatalf("Could not find item with ID '%s'", itemID)
// 	}

// 	dataItem := Data{}

// 	err = dynamodbattribute.UnmarshalMap(result.Item, &dataItem)
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
// 	}

// 	fmt.Println("Found Data Item:")
// 	fmt.Println("ID:          ", dataItem.ID)
// 	fmt.Println("Temperature: ", dataItem.Temperature)
// 	fmt.Println("Llm:         ", dataItem.Llm)
// 	fmt.Println("Rookie:      ", dataItem.Rookie)
// 	fmt.Println("Prompt:      ", dataItem.Prompt)
// 	fmt.Println("Context:     ", dataItem.Context)
// 	fmt.Println("Process:     ", dataItem.Process)
// }

// func readMovieItem(svc *dynamodb.DynamoDB) {
// 	tableName := "Movies"
// 	movieName := "The Big New Movie"
// 	movieYear := "2015"

// 	result, err := svc.GetItem(&dynamodb.GetItemInput{
// 		TableName: aws.String(tableName),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"Year": {
// 				N: aws.String(movieYear),
// 			},
// 			"Title": {
// 				S: aws.String(movieName),
// 			},
// 		},
// 	})
// 	if err != nil {
// 		log.Fatalf("Got error calling GetMovie: %s", err)
// 	}

// 	if result.Item == nil {
// 		log.Fatalf("Could not find '%s'", movieName)
// 	}

// 	Movie := Movie{}

// 	err = dynamodbattribute.UnmarshalMap(result.Item, &Movie)
// 	if err != nil {
// 		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
// 	}

// 	fmt.Println("Found Movie:")
// 	fmt.Println("Year:  ", Movie.Year)
// 	fmt.Println("Title: ", Movie.Title)
// 	fmt.Println("Plot:  ", Movie.Plot)
// 	fmt.Println("Rating:", Movie.Rating)
// }

// func listTables(svc *dynamodb.DynamoDB) {
// 	result, err := svc.ListTables(&dynamodb.ListTablesInput{})
// 	if err != nil {
// 		log.Panicf("error: %v", err)
// 	}

// 	fmt.Println("DynamoDB tables:")
// 	for _, table := range result.TableNames {
// 		fmt.Println(*table)
// 	}
// }

// func dynamodbConfig() *dynamodb.DynamoDB {
// 	config := &aws.Config{
// 		Endpoint: aws.String("http://dynamodb-local:5555"), // Docker Compose service name and port
// 		Region:   aws.String("local"),                      // Use 'local' for local DynamoDB
// 		Credentials: credentials.NewStaticCredentials(
// 			"fakeAccessKey",
// 			"fakeSecretKey",
// 			"",
// 		),
// 	}

// 	// Create a new session
// 	sess, err := session.NewSession(config)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Create a DynamoDB client
// 	svc := dynamodb.New(sess)

// 	return svc
// }

// func dynamodbConfig2() *dynamodb.DynamoDB {
// 	// Initialize a session that the SDK will use to load
// 	// credentials from the shared credentials file ~/.aws/credentials
// 	// and region from the shared configuration file ~/.aws/config.
// 	sess := session.Must(session.NewSessionWithOptions(session.Options{
// 		SharedConfigState: session.SharedConfigEnable,
// 	}))

// 	// Create DynamoDB client
// 	svc := dynamodb.New(sess)

// 	return svc
// }

// func listTables2(svc *dynamodb.DynamoDB) {
// 	// create the input configuration instance
// 	input := &dynamodb.ListTablesInput{}

// 	fmt.Printf("Tables:\n")

// 	for {
// 		// Get the list of tables
// 		result, err := svc.ListTables(input)
// 		if err != nil {
// 			if aerr, ok := err.(awserr.Error); ok {
// 				switch aerr.Code() {
// 				case dynamodb.ErrCodeInternalServerError:
// 					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
// 				default:
// 					fmt.Println(aerr.Error())
// 				}
// 			} else {
// 				// Print the error, cast err to awserr.Error to get the Code and
// 				// Message from an error.
// 				fmt.Println(err.Error())
// 			}
// 			return
// 		}

// 		for _, n := range result.TableNames {
// 			fmt.Println(*n)
// 		}

// 		// assign the last read tablename as the start for our next call to the ListTables function
// 		// the maximum number of table names returned in a call is 100 (default), which requires us to make
// 		// multiple calls to the ListTables function to retrieve all table names
// 		input.ExclusiveStartTableName = result.LastEvaluatedTableName

// 		if result.LastEvaluatedTableName == nil {
// 			break
// 		}
// 	}
// }

// func scanDataItem(svc *dynamodb.DynamoDB) {
// 	tableName := "Data2"
// 	minTemperature := 25

// 	// Crear la expresión para llenar la estructura de entrada
// 	// Obtener todos los elementos con temperatura superior a la mínima especificada
// 	filt := expression.Name("Temperature").GreaterThanEqual(expression.Value(minTemperature))

// 	// Obtener de vuelta el ID y la temperatura
// 	proj := expression.NamesList(expression.Name("ID"), expression.Name("Temperature"))

// 	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
// 	if err != nil {
// 		log.Fatalf("Got error building expression: %s", err)
// 	}

// 	// Construir los parámetros de entrada para la consulta
// 	params := &dynamodb.ScanInput{
// 		ExpressionAttributeNames:  expr.Names(),
// 		ExpressionAttributeValues: expr.Values(),
// 		FilterExpression:          expr.Filter(),
// 		ProjectionExpression:      expr.Projection(),
// 		TableName:                 aws.String(tableName),
// 	}

// 	// Hacer la llamada a la API de Scan de DynamoDB
// 	result, err := svc.Scan(params)
// 	if err != nil {
// 		log.Fatalf("Query API call failed: %s", err)
// 	}

// 	numItems := 0

// 	for _, i := range result.Items {
// 		item := Data{}

// 		err = dynamodbattribute.UnmarshalMap(i, &item)

// 		if err != nil {
// 			log.Fatalf("Got error unmarshalling: %s", err)
// 		}

// 		numItems++

// 		fmt.Println("ID: ", item.ID)
// 		fmt.Println("Temperature: ", item.Temperature)
// 		fmt.Println("Llm: ", item.Llm)
// 		fmt.Println("Rookie: ", item.Rookie)
// 		fmt.Println("Prompt: ", item.Prompt)
// 		fmt.Println("Context: ", item.Context)
// 		fmt.Println("Process: ", item.Process)
// 		fmt.Println("Date From: ", item.DateFrom)
// 		fmt.Println("Date To: ", item.DateTo)
// 		fmt.Println()
// 	}

// 	fmt.Println("Found", numItems, "item(s) with a temperature above or equal to", minTemperature)
// }

// func scanMovieItem(svc *dynamodb.DynamoDB) {
// 	tableName := "Movies"
// 	minRating := 4.0
// 	year := 2013

// 	// Create the Expression to fill the input struct with.
// 	// Get all movies in that year; we'll pull out those with a higher rating later
// 	filt := expression.Name("Year").Equal(expression.Value(year))

// 	// Or we could get by ratings and pull out those with the right year later
// 	//    filt := expression.Name("info.rating").GreaterThan(expression.Value(min_rating))

// 	// Get back the title, year, and rating
// 	proj := expression.NamesList(expression.Name("Title"), expression.Name("Year"), expression.Name("Rating"))

// 	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
// 	if err != nil {
// 		log.Fatalf("Got error building expression: %s", err)
// 	}

// 	// Build the query input parameters
// 	params := &dynamodb.ScanInput{
// 		ExpressionAttributeNames:  expr.Names(),
// 		ExpressionAttributeValues: expr.Values(),
// 		FilterExpression:          expr.Filter(),
// 		ProjectionExpression:      expr.Projection(),
// 		TableName:                 aws.String(tableName),
// 	}

// 	// Make the DynamoDB Query API call
// 	result, err := svc.Scan(params)
// 	if err != nil {
// 		log.Fatalf("Query API call failed: %s", err)
// 	}

// 	numItems := 0

// 	for _, i := range result.Items {
// 		item := Movie{}

// 		err = dynamodbattribute.UnmarshalMap(i, &item)

// 		if err != nil {
// 			log.Fatalf("Got error unmarshalling: %s", err)
// 		}

// 		// Which ones had a higher rating than minimum?
// 		if item.Rating > minRating {
// 			// Or it we had filtered by rating previously:
// 			//   if item.Year == year {
// 			numItems++

// 			fmt.Println("Title: ", item.Title)
// 			fmt.Println("Rating:", item.Rating)
// 			fmt.Println()
// 		}
// 	}

// 	fmt.Println("Found", numItems, "movie(s) with a rating above", minRating, "in", year)
// }

// func deleteMovieItem(svc *dynamodb.DynamoDB) {
// 	tableName := "Movies"
// 	movieName := "The Big New Movie"
// 	movieYear := "2015"

// 	input := &dynamodb.DeleteItemInput{
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"Year": {
// 				N: aws.String(movieYear),
// 			},
// 			"Title": {
// 				S: aws.String(movieName),
// 			},
// 		},
// 		TableName: aws.String(tableName),
// 	}

// 	_, err := svc.DeleteItem(input)
// 	if err != nil {
// 		log.Fatalf("Got error calling DeleteItem: %s", err)
// 	}

// 	fmt.Println("Deleted '" + movieName + "' (" + movieYear + ") from table " + tableName)
// }

// func deleteDataItem(svc *dynamodb.DynamoDB) {
// 	tableName := "Data2"
// 	itemID := "12345"

// 	input := &dynamodb.DeleteItemInput{
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"ID": {
// 				S: aws.String(itemID),
// 			},
// 		},
// 		TableName: aws.String(tableName),
// 	}

// 	_, err := svc.DeleteItem(input)
// 	if err != nil {
// 		log.Fatalf("Got error calling DeleteItem: %s", err)
// 	}

// 	fmt.Println("Deleted item with ID '" + itemID + "' from table " + tableName)
// }

// func updateMovieItem(svc *dynamodb.DynamoDB) {
// 	tableName := "Movies"
// 	movieName := "The Big New Movie"
// 	movieYear := "2015"
// 	movieRating := "0.5"

// 	input := &dynamodb.UpdateItemInput{
// 		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
// 			":r": {
// 				N: aws.String(movieRating),
// 			},
// 		},
// 		TableName: aws.String(tableName),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"Year": {
// 				N: aws.String(movieYear),
// 			},
// 			"Title": {
// 				S: aws.String(movieName),
// 			},
// 		},
// 		ReturnValues:     aws.String("UPDATED_NEW"),
// 		UpdateExpression: aws.String("set Rating = :r"),
// 	}

// 	_, err := svc.UpdateItem(input)
// 	if err != nil {
// 		log.Fatalf("Got error calling UpdateItem: %s", err)
// 	}

// 	fmt.Println("Successfully updated '" + movieName + "' (" + movieYear + ") rating to " + movieRating)
// }

// func updateDataItem(svc *dynamodb.DynamoDB) {
// 	tableName := "Data2"
// 	itemID := "12345"
// 	newTemperature := "30"

// 	input := &dynamodb.UpdateItemInput{
// 		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
// 			":t": {
// 				N: aws.String(newTemperature),
// 			},
// 		},
// 		TableName: aws.String(tableName),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"ID": {
// 				S: aws.String(itemID),
// 			},
// 		},
// 		ReturnValues:     aws.String("UPDATED_NEW"),
// 		UpdateExpression: aws.String("set Temperature = :t"),
// 	}

// 	_, err := svc.UpdateItem(input)
// 	if err != nil {
// 		log.Fatalf("Got error calling UpdateItem: %s", err)
// 	}

// 	fmt.Println("Successfully updated item with ID '" + itemID + "' temperature to " + newTemperature)
// }

// // Get table items from JSON file
// func getItemsFromJson[T any](filename string) []T {
// 	raw, err := os.ReadFile(filename)
// 	if err != nil {
// 		log.Fatalf("Got error reading file: %s", err)
// 	}

// 	var items []T
// 	err = json.Unmarshal(raw, &items)
// 	if err != nil {
// 		log.Fatalf("Got error unmarshalling JSON: %s", err)
// 	}

// 	return items
// }

// // Example using CloudQuery type
// func loadDataItemsJson(svc *dynamodb.DynamoDB) {
// 	// Get table items from .data.json
// 	cqItems := getItemsFromJson[Data]("data_items.json")

// 	// Add each item to Data2 table:
// 	tableName := "Data2"

// 	for _, item := range cqItems {
// 		av, err := dynamodbattribute.MarshalMap(item)
// 		if err != nil {
// 			log.Fatalf("Got error marshalling map: %s", err)
// 		}

// 		// Create item in table Data2
// 		input := &dynamodb.PutItemInput{
// 			Item:      av,
// 			TableName: aws.String(tableName),
// 		}

// 		_, err = svc.PutItem(input)
// 		if err != nil {
// 			log.Fatalf("Got error calling PutItem: %s", err)
// 		}

// 		// Assuming cq type has an 'ID' field
// 		fmt.Println("Successfully added item with ID '" + item.ID + "' to table " + tableName)
// 	}
// }

// func loadMovieItemsJson(svc *dynamodb.DynamoDB) {
// 	// Get table items from .movie_data.json
// 	items := getItemsFromJson[Movie]("./movie_data.json")

// 	// Add each item to Movies table:
// 	tableName := "Movies"

// 	for _, item := range items {
// 		av, err := dynamodbattribute.MarshalMap(item)
// 		if err != nil {
// 			log.Fatalf("Got error marshalling map: %s", err)
// 		}

// 		// Create item in table Movies
// 		input := &dynamodb.PutItemInput{
// 			Item:      av,
// 			TableName: aws.String(tableName),
// 		}

// 		_, err = svc.PutItem(input)
// 		if err != nil {
// 			log.Fatalf("Got error calling PutItem: %s", err)
// 		}

// 		year := strconv.Itoa(item.Year)

// 		fmt.Println("Successfully added '" + item.Title + "' (" + year + ") to table " + tableName)
// 	}
// }
