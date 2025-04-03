
func handler(w http.ResponseWriter, r  *http.Request) {
	switch r.Method {
		case "GET":
			fmt.Fprintf(w, "HERE GET\n")
		case "POST":
			fmt.Fprintf(w, "HERE POST\n")
		case "PATCH":
			fmt.Fprintf(w, "HERE PATCH\n")
		default:
			fmt.Fprintf(w, "Sorry, only GET, POST and PATCH methods are supported\n")
	}
}

func server() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	loggedMux := output.Logger(mux)

	fmt.Println("Starting server on :8000...")
	log.Fatal(http.ListenAndServe(":8000", loggedMux))
}
