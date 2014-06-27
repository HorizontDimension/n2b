'use strict';

/* Directives */


angular.module('myApp.directives', []).
  directive('appVersion', ['version', function(version) {
    return function(scope, elm, attrs) {
      elm.text(version);
    };
  }]).directive('fixOverflow', [function () {
    return {
        restrict: 'A',
        scope: true,
       link: function (scope, elem, attrs, control,window) {


        }
    };
}])
