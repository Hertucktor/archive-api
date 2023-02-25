package apiv1

import (
	"github.com/Hertucktor/archive-api/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type SetNameInfo struct {
	SetName  []string `json:"setName" bson:"setName"`
	Created  string   `json:"created" bson:"created"`
	Modified string   `json:"modified" bson:"modified"`
}

func SetupRoutes(logger *zap.SugaredLogger, port string) {
	logger.Infow("Starting API on port %v", port)

	//Create Router
	router := gin.Default()
	router.NoRoute(utils.NotFoundHandler)

	// V1
	//v1 := router.Group("/v1")
	//Status calls
	statusRouter := router.Group("/status")
	{
		statusRouter.GET("/alive", statusAlive)
		statusRouter.GET("/check", statusCheck)
	}
	//CRUD route for card info
	//api := router.PathPrefix("/api").Subrouter()
	//api.Use(mux.CORSMethodMiddleware(api), corsOriginMiddleware)
	//api.HandleFunc("/card", createNewCardEntryOnAllCardCollection).Methods(http.MethodPost)
	//api.HandleFunc("/cards", returnAllCardEntriesFromAllCardCollection).Methods(http.MethodGet)                            //From allCards Coll
	//api.HandleFunc("/cards/set-names/{setName}", returnAllCardsBySetFromAllCardCollection).Methods(http.MethodGet)         //From allCards Coll
	//api.HandleFunc("/cards/collector-numbers/{number}/set-names/{setName}", readFromOwnCollection).Methods(http.MethodGet) //From myCards Coll
	//api.HandleFunc("/cards/collector-number/{number}/set-names/{setName}", updateSingleCardFromOwnCollection).Methods(http.MethodPut)
	//api.HandleFunc("/cards/collector-number/{number}/set-names/{setName}", deleteSingleCardFromOwnCollection).Methods(http.MethodDelete)
	//api.HandleFunc("/cards/set-names", returnAllSetName).Methods(http.MethodGet)
	//API Operations for img info
	//img := router.PathPrefix("/img").Subrouter()
	//img.HandleFunc("/set-names/{setName}", returnSingleImg).Methods(http.MethodGet)

	//Routes for Image manipulation
	//upload := router.PathPrefix("/uploads").Subrouter()
	//upload.HandleFunc("/img", uploadImg).Methods(http.MethodPost)

	//Serve UI
	//staticDir := "/static/"
	//router.PathPrefix(staticDir).Handler(http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	//Open http connection
	if err := http.ListenAndServe(port, router); err != nil {
		logger.Errorw("Panic: problem with TCP network connection", err)
	}
}

//
//func returnAllSetName(w http.ResponseWriter, r *http.Request) {
//	log.Info().Msg("Endpoint Hit: returnAllCardsBySetFromAllCardCollection")
//
//	conf, err := config.GetConfig("config.yml")
//	if err != nil {
//		log.Fatal().Timestamp().Err(err).Msg("Fatal: couldn't receive env vars")
//	}
//
//	client, ctx, err := buildClient(conf)
//	if err != nil {
//		log.Fatal().Timestamp().Err(err).Msg("v: couldn't build client")
//	}
//
//	setNameInfo := allSetNames(client, ctx, conf)
//	if err != nil {
//		w.WriteHeader(500)
//		_, _ = w.Write([]byte("set names couldn't been found"))
//		log.Error().Timestamp().Err(err).Msg("Error: couldn't return set names")
//		return
//	}
//
//	setNamesBytes, err := json.Marshal(setNameInfo)
//	if err != nil {
//		log.Error().Timestamp().Err(err)
//		w.WriteHeader(500)
//		return
//	}
//
//	if _, err = w.Write(setNamesBytes); err != nil {
//		log.Error().Timestamp().Err(err)
//		w.WriteHeader(500)
//		return
//	}
//	w.WriteHeader(200)
//
//}
//
//func allSetNames(client *mongo.Client, ctx context.Context, conf config.Config) SetNameInfo {
//	var setNameInfo SetNameInfo
//	collection := client.Database(conf.DBName).Collection("setNames")
//
//	result := collection.FindOne(ctx, bson.M{}).Decode(&setNameInfo)
//	fmt.Println(result)
//	fmt.Println(setNameInfo)
//
//	return setNameInfo
//}
//
//func buildClient(conf config.Config) (*mongo.Client, context.Context, error) {
//
//	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + conf.DBUser + ":" + conf.DBPass + "@" + conf.DBPort + "/" + conf.DBName))
//	if err != nil {
//		log.Error().Err(err)
//		return client, nil, err
//	}
//
//	ctx := context.TODO()
//	if err = client.Connect(ctx); err != nil {
//		log.Error().Err(err)
//		return client, ctx, err
//	}
//
//	return client, ctx, err
//}