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
}]).directive('qnValidate', [
        function() {
            return {
                link: function(scope, element, attr) {
                    var form = element.inheritedData('$formController');
                    // no need to validate if form doesn't exists
                    if (!form) return;
                    // validation model
                    var validate = attr.qnValidate;
                    // watch validate changes to display validation
                    scope.$watch(validate, function(errors) {

                        // every server validation should reset others
                        // note that this is form level and NOT field level validation
                        form.$serverError = { };

                        // if errors is undefined or null just set invalid to false and return
                        if (!errors) {
                            form.$serverInvalid = false;
                            return;
                        }
                        // set $serverInvalid to true|false
                        form.$serverInvalid = (errors.length > 0);

                        // loop through errors
                        angular.forEach(errors, function(error, i) {
                            form.$serverError[error.key] = { $invalid: true, message: error.value };
                        });
                    });
                }
            };
        }
    ]).directive('validFile',function(){
    return {
        require:'ngModel',
        link:function(scope,el,attrs,ngModel){
            el.bind('change',function(){
                scope.$apply(function(){
                    ngModel.$setViewValue(el.val());
                    ngModel.$render();
                });
            });
        }
    }
});
