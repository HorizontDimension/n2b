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
    }).factory('Contacts', function($resource){
        var resource = $resource('http://localhost:8080/contacts/new',{},{
            new:{
                method:"POST",
                isArray:true
               // headers:{'Content-Type':'application/x-www-form-urlencoded; charset=UTF-8'}
            }
        });
        return resource;
    });
