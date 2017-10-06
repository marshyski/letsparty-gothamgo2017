# letsparty-gothamgo2017

Talk presentation for "Building a Distributed Serverless Platform from Scratch" at Gotham Go 2017 - https://drive.google.com/open?id=12vFoiN1llJqdX9XbPS2FxSX0XmnEHzS0O4P05HM_fRU

    go test -bench=. -timeout=0 > out
    cat out | benchgraph -title='sync.Map and RWMutex Stable Keys Load, Store, LoadStore, Delete' -obn="CMGet,SMLoad,CMCreate,SMStore,CMGetCreate,SMLoadStore,CMDelete,SMDelete" -oba="1,2,4,8,16,32,64"
    
 **Influenced by https://github.com/deckarep/sync-map-analysis**
