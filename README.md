ElkAPI
======

ElasticSearch restful API, include ElkInput(), ElkGetAll()  <br>

ELkInput(elkurl, index, type string, obj interface{})   <br>
       input:   <br>
       ---	elkurl string   //ElasticSearch Server Destination, ex: http://g2tool.cloudapp.net:9200/   <br>
       ---    index string    //ElasticSearch index    <br>
       ---    type string    //ElasticSearch type    <br>
       ---    obj interface{} //Structure Object that want to put in ElasticSearch  <br>

<br>

ElkGetAll(index, type string)  <br>
       input: <br>
       ---    index string    //ElasticSearch index  <br>
       ---    type string    //ElasticSearch type  <br>
