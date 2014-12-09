ElkAPI
======

ElasticSearch restful API, include ElkInput(), ElkGetAll()  <br>

ELkInput(elkurl, index, table string, obj interface{})   <br>
       input:   <br>
       ---	elkurl string   //ElasticSearch Server Destination, ex: http://g2tooles2.cloudapp.net:9200/   <br>
       ---    index string    //ElasticSearch index    <br>
       ---    table string    //ElasticSearch table    <br>
       ---    obj interface{} //Structure Object that want to put in ElasticSearch  <br>

<br>

ElkGetAll(index, table string)  <br>
       input: <br>
       ---    index string    //ElasticSearch index  <br>
       ---    table string    //ElasticSearch table  <br>
