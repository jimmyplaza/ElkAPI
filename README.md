ElkAPI
======

ElasticSearch restful API, include ElkInput(), ElkGetAll()  <br>

ELkInput(elkurl, index, es_type string, obj interface{})   <br>
       input:   <br>
       ---	elkurl string   //ElasticSearch Server Destination, ex: http://g2tooles2.cloudapp.net:9200/   <br>
       ---    index string    //ElasticSearch index    <br>
       ---    es_type string    //ElasticSearch es_type    <br>
       ---    obj interface{} //Structure Object that want to put in ElasticSearch  <br>

<br>

ElkGetAll(index, es_type string)  <br>
       input: <br>
       ---    index string    //ElasticSearch index  <br>
       ---    es_type string    //ElasticSearch es_type  <br>
