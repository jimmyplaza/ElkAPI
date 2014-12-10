ElkAPI
======

ElasticSearch restful API, include ElkInput(), ElkGetAll()  <br>


* Warning: index, type name must be lowercase letter, also the source json content must be lowercase letter. 


ELkInput(elkurl, index, es_type string, obj interface{})   <br>
=======
       input:   <br>
       ---	elkurl string   //ElasticSearch Server Destination, ex: http://g2tool.cloudapp.net:9200/   <br>
       ---    index string    //ElasticSearch index    <br>
       ---    es_type string    //ElasticSearch es_type    <br>
       ---    obj interface{} //Structure Object that want to put in ElasticSearch  <br>

<br>
ElkGetAll(index, es_type string)  <br>
       input: <br>
       ---    index string    //ElasticSearch index  <br>
       ---    es_type string    //ElasticSearch es_type  <br>
