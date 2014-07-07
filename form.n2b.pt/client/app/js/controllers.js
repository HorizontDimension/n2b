'use strict';

/* Controllers */

angular.module('myApp.controllers', [])
    .controller('upgrade', ['$scope', '$fileUploader', function ($scope,$fileUploader) {
        $scope.submitted = false;

        $scope.upgrade = {};

        $scope.success = "";

        // create a uploader with options
        var uploader = $scope.uploader = $fileUploader.create({
            scope: $scope,                          // to automatically update the html. Default: $rootScope
            url: 'http://localhost:8080/upgrades/new',
            formData: [ $scope.upgrade],
            queueLimit: 1
        });


        // ADDING FILTERS


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
            item.formData.push({challenge: $scope.transfer.captcha.challenge});
            item.formData.push({response: $scope.transfer.captcha.response});
            delete item.formData.captcha;
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
            $scope.errors = response;
        });

        uploader.bind('complete', function (event, xhr, item, response) {
            console.info('Complete', xhr, item, response);
            $scope.errors = response;

            if (response == "") {
                $scope.success = "sucesso"
            }
        });

        uploader.bind('progressall', function (event, progress) {
            console.info('Total progress: ' + progress);
        });

        uploader.bind('completeall', function (event, items) {
            console.info('Complete all', items);
        });


        $scope.submit = function () {

            if ($scope.Upgrade.$valid) {
                // Submit as normal
                console.log($scope.transfer)
                uploader.uploadAll()
            }
            else {
                console.log("we got errors")
                console.log($scope.Transfer)
                $scope.submitted = true;
            }
        }
    }])

    .controller('acerca', ['$scope', 'Transfers', function ($scope) {

    }])

    .controller('contacts', ['$scope', 'Contacts', function ($scope,Contacts) {
        $scope.submitted = false;


        $scope.contact = {};

        $scope.success = "";

        $scope.submit= function(){
            if ($scope.Contact.$valid) {
                Contacts.new({},$scope.contact,
                    function(response, headers) {
                        $scope.errors = response;
                }, function(errorResponse) {
                        $scope.errors = errorResponse;
                } )
            }
            else {
                $scope.submitted = true;
            }
        }
    }])

    .controller('menu', ['$scope', '$location', function ($scope, $location) {

        $scope.getClass = function (path) {
            if ($location.path().substr(0, path.length) == path) {
                return "active"
            } else {
                return ""
            }
        }

    }])
    .controller('login', ['$scope', 'Transfers', function ($scope) {

    }])
    .controller('transfer', ['$scope', 'Transfers', '$fileUploader', function ($scope, Transfers, $fileUploader) {
        $scope.submitted = false;


        $scope.transfer = {};
        $scope.transfer.captcha = {}
        $scope.success = "";

        // create a uploader with options
        var uploader = $scope.uploader = $fileUploader.create({
            scope: $scope,                          // to automatically update the html. Default: $rootScope
            url: 'http://localhost:8080/transfers/new',
            formData: [ $scope.transfer],
            queueLimit: 1
        });


        // ADDING FILTERS


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
            item.formData.push({challenge: $scope.transfer.captcha.challenge});
            item.formData.push({response: $scope.transfer.captcha.response});
            delete item.formData.captcha;
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
            $scope.errors = response;
        });

        uploader.bind('complete', function (event, xhr, item, response) {
            console.info('Complete', xhr, item, response);
            $scope.errors = response;

            if (response == "") {
                $scope.success = "sucesso"
            }
        });

        uploader.bind('progressall', function (event, progress) {
            console.info('Total progress: ' + progress);
        });

        uploader.bind('completeall', function (event, items) {
            console.info('Complete all', items);
        });


        $scope.submit = function () {

            if ($scope.Transfer.$valid) {
                // Submit as normal
                console.log($scope.transfer)
                uploader.uploadAll()
            }
            else {
                console.log("we got errors")
                console.log($scope.Transfer)
                $scope.submitted = true;
            }
        }


    }]);
