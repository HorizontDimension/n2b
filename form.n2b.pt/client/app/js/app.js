'use strict';


// Declare app level module which depends on filters, and services
angular.module('myApp', [
  'ngRoute',
  'angularFileUpload',
  'ngResource',
  'vcRecaptcha',
  'myApp.filters',
  'myApp.services',
  'myApp.directives',
  'myApp.controllers'
]).
config(['$routeProvider', function($routeProvider) {
    $routeProvider.when('/acerca', {templateUrl: '/partials/acerca.html', controller: 'acerca'});
    $routeProvider.when('/contactos', {templateUrl: '/partials/contacts.html', controller: 'contacts'});
    $routeProvider.when('/upgrade', {templateUrl: '/partials/upgrade.html', controller: 'MyCtrl1'});
    $routeProvider.when('/transfer', {templateUrl: '/partials/transfer.html', controller: 'MyCtrl2'});
    $routeProvider.when('/login', {templateUrl: '/partials/login.html', controller: 'login'});
    $routeProvider.otherwise({redirectTo: '/acerca'});
}]);
