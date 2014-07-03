'use strict';

/* Services */


// Demonstrate how to register services
// In this case it is a simple value service.
angular.module('myApp.services', []).
  value('version', '0.1').factory('Transfers', function($resource){
        var resource = $resource('http://n2b.go.euroneves.pt/transfers/new',{},{
            new:{
                method:"POST",
                isArray:false,
                headers:{'Content-Type':'application/x-www-form-urlencoded; charset=UTF-8'}
            }
        });
        return resource;
    });
