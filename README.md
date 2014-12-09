ElkAPI
======

ElasticSearch restful api, include ElkInput(), ElkGetAll()  <br>

ELkInput(elkurl, index, table string, obj interface{})   <br>
       input:   <br>
       &nbsp	elkurl string   //ElasticSearch Server Destination, ex: http://g2tooles2.cloudapp.net:9200/   <br>
       &nbsp  index string    //ElasticSearch index    <br>
       &nbsp  table string    //ElasticSearch table    <br>
       &nbsp  obj interface{} //Structure Object that want to put in ElasticSearch  <br>



ElkGetAll(index, table string)  <br>
       input: <br>
              index string    //ElasticSearch index  <br>
              table string    //ElasticSearch table  <br>
