'use strict';

/* Controllers */

angular.module('myApp.controllers', [])
    .controller('MyCtrl1', ['$scope', 'Transfers', function ($scope) {

    }])
    .controller('MyCtrl2', ['$scope','Transfers',  function ($scope, Transfers) {
        $scope.agent = {};



        $scope.Submit = function (agent) {
            console.log("teste")
            Transfers.new({},
                jQuery.param({
                "OldName": "asdasd",
                "NewName": "asdasd",
                "OldNif": "asdasd",
                "NewNif": "asdasd",
                "Hardlock": "asdasd",
                "Proof": "asdasd"
            }))


        }


    }]);
