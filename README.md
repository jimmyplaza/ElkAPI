ElkAPI
======

ElasticSearch restful api, include ElkInput(), ElkGetAll()


ELkInput()
input: elkurl string   //ElasticSearch Server Destination, ex: http://g2tooles2.cloudapp.net:9200/
       index string    //ElasticSearch index
       table string    //ElasticSearch table
       obj interface{}) //Structure Object that want to put in ElasticSearch


ElkGetAll()
input: 
      index string    //ElasticSearch index
      table string    //ElasticSearch table
