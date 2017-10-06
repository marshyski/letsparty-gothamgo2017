# letsparty-gothamgo2017


    go test -bench=. -timeout=0 > out
    cat out | benchgraph -title='sync.Map and RWMutex Stable Keys Load, Store, LoadStore, Delete' -obn="CMGet,SMLoad,CMCreate,SMStore,CMGetCreate,SMLoadStore,CMDelete,SMDelete" -oba="1,2,4,8,16,32,64"
    
 **Influenced by https://github.com/deckarep/sync-map-analysis**
