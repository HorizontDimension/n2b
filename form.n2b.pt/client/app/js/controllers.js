'use strict';

/* Controllers */

angular.module('myApp.controllers', [])
    .controller('MyCtrl1', ['$scope', 'Transfers', function ($scope) {

    }])
    .controller('acerca', ['$scope', 'Transfers', function ($scope) {

    }])
    .controller('contacts', ['$scope', 'Transfers', function ($scope) {

    }])
    .controller('menu', ['$scope', '$location', function ($scope,$location) {

        $scope.getClass = function(path) {
            if ($location.path().substr(0, path.length) == path) {
                return "active"
            } else {
                return ""
            }
        }

    }])
    .controller('login', ['$scope', 'Transfers', function ($scope) {

    }])
    .controller('MyCtrl2', ['$scope','Transfers','$fileUploader',  function ($scope, Transfers,$fileUploader) {
        $scope.submitted = false;



        $scope.agent = {};
        $scope.success= "";



        // create a uploader with options
        var uploader = $scope.uploader = $fileUploader.create({
            scope: $scope,                          // to automatically update the html. Default: $rootScope
            url: 'http://localhost:8080/transfers/new',
            formData:[ $scope.agent],
            queueLimit:1,
            filters: [
                function (item) {                    // first user filter
                    console.info('filter1');
                    return true;
                }
            ]
        });





        // ADDING FILTERS

        uploader.filters.push(function (item) { // second user filter
            console.info('filter2');
            return true;
        });

        // REGISTER HANDLERS

        uploader.bind('afteraddingfile', function (event, item) {
            console.info('After adding a file', item);
        });

        uploader.bind('whenaddingfilefailed', function (event, item) {
            console.info('When adding a file failed', item);
        });

        uploader.bind('afteraddingall', function (event, items) {
            console.info('After adding all files', items);
        });

        uploader.bind('beforeupload', function (event, item) {
            console.info('Before upload', item);
        });

        uploader.bind('progress', function (event, item, progress) {
            console.info('Progress: ' + progress, item);
        });

        uploader.bind('success', function (event, xhr, item, response) {
            console.info('Success', xhr, item, response);
        });

        uploader.bind('cancel', function (event, xhr, item) {
            console.info('Cancel', xhr, item);
        });

        uploader.bind('error', function (event, xhr, item, response) {
            console.info('Error', xhr, item, response);
            $scope.errors=response;
        });

        uploader.bind('complete', function (event, xhr, item, response) {
            console.info('Complete', xhr, item, response);
            $scope.errors=response;
           if (response==""){
               $scope.success="sucesso"
           }


        });

        uploader.bind('progressall', function (event, progress) {
            console.info('Total progress: ' + progress);
        });

        uploader.bind('completeall', function (event, items) {
            console.info('Complete all', items);
        });


        $scope.submit=function(){
            console.log("submited")

            if ($scope.Transfer.$valid) {
                // Submit as normal
               $scope.uploader.uploadAll()
            }
            else {
                console.log("we got errors")
                console.log($scope.Transfer)
                $scope.submitted = true;
            }
        }








    }]);
