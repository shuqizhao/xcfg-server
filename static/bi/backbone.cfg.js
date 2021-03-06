﻿define(function(require) {
    require('jquery');
    require('cookie');
    require('bootstarp');
    var Backbone = require('backbone');

    (function($) {
        $.fn.extend({
            "getQueryObject": function() {
                // get page location from variable if history fragment is empty
                var fragmentStr = Backbone.history.getFragment();
                var queryObject = {};
                var markIndex = fragmentStr.indexOf("?");

                // contain query string in fragments
                if (markIndex > -1) {
                    // build query object
                    var queryString = fragmentStr.slice(markIndex + 1);
                    if (queryString.indexOf("url=") >= 0) {
                        var queryArray = queryString.split('url=');
                        queryObject["url"] = decodeURIComponent(queryArray[1]);
                    } else {
                        var queryArray = queryString.split('&');
                        for (var i = 0; i < queryArray.length; i++) {
                            var queryPair = queryArray[i].split('=');
                            queryObject[queryPair[0]] = decodeURIComponent(queryPair[1]);
                        }
                    }

                }
                return queryObject;
            },
            "getRootRouter": function() {
                // get page location from variable if history fragment is empty
                var fragmentStr = Backbone.history.getFragment();
                var markIndex = fragmentStr.indexOf("/");

                // contain query string in fragments
                if (markIndex > -1) {
                    // build query object
                    return fragmentStr.substring(0, markIndex);
                }
                return "";
            },
            "navigate": function() {
                var url = this.getRootRouter();
                if (url == '') {
                    history.go(-1);
                } else {
                    //change location and jump
                    Backbone.history.navigate(url);
                    //change location and not jump
                    //Backbone.history.navigate(url,{trigger: true,replace: true});
                    var Module = require('/bi/biz/' + url + '/' + url + '.js');
                    var m = new Module({
                        el: $('.showbox')
                    });
                    m.render();
                }
            },
            "validateXML": function(xmlContent) {
                //errorCode 0是xml正确，1是xml错误，2是无法验证 
                var xmlDoc, errorMessage, errorCode = 0;
                // code for IE 
                if (window.ActiveXObject) {
                    xmlDoc = new ActiveXObject("Microsoft.XMLDOM");
                    xmlDoc.async = "false";
                    xmlDoc.loadXML(xmlContent);

                    if (xmlDoc.parseError.errorCode != 0) {
                        errorMessage = "错误code: " + xmlDoc.parseError.errorCode + "\n";
                        errorMessage = errorMessage + "错误原因: " + xmlDoc.parseError.reason;
                        errorMessage = errorMessage + "错误位置: " + xmlDoc.parseError.line;
                        errorCode = 1;
                    } else {
                        errorMessage = "格式正确";
                    }
                }
                // code for Mozilla, Firefox, Opera, chrome, safari,etc. 
                else if (document.implementation.createDocument) {
                    var parser = new DOMParser();
                    xmlDoc = parser.parseFromString(xmlContent, "text/xml");
                    var error = xmlDoc.getElementsByTagName("parsererror");
                    if (error.length > 0) {
                        if (xmlDoc.documentElement.nodeName == "parsererror") {
                            errorCode = 1;
                            errorMessage = xmlDoc.documentElement.childNodes[0].nodeValue;
                        } else {
                            errorCode = 1;
                            errorMessage = xmlDoc.getElementsByTagName("parsererror")[0].innerHTML;
                        }
                    } else {
                        errorMessage = "格式正确";
                    }
                } else {
                    errorCode = 2;
                    errorMessage = "浏览器不支持验证，无法验证xml正确性";
                }
                return {
                    "msg": errorMessage,
                    "error_code": errorCode
                };
            }
        });
        return this;
    })(jQuery);
    (function($, h, c) {
        var a = $([]),
            e = $.resize = $.extend($.resize, {}),
            i, k = "setTimeout",
            j = "resize",
            d = j + "-special-event",
            b = "delay",
            f = "throttleWindow";
        e[b] = 250;
        e[f] = true;
        $.event.special[j] = {
            setup: function() {
                if (!e[f] && this[k]) { return false }
                var l = $(this);
                a = a.add(l);
                $.data(this, d, { w: l.width(), h: l.height() });
                if (a.length === 1) { g() }
            },
            teardown: function() {
                if (!e[f] && this[k]) { return false }
                var l = $(this);
                a = a.not(l);
                l.removeData(d);
                if (!a.length) { clearTimeout(i) }
            },
            add: function(l) {
                if (!e[f] && this[k]) { return false }
                var n;

                function m(s, o, p) {
                    var q = $(this),
                        r = $.data(this, d);
                    r.w = o !== c ? o : q.width();
                    r.h = p !== c ? p : q.height();
                    n.apply(this, arguments)
                }
                if ($.isFunction(l)) { n = l; return m } else {
                    n = l.handler;
                    l.handler = m
                }
            }
        };

        function g() {
            i = h[k](function() {
                a.each(function() {
                    var n = $(this),
                        m = n.width(),
                        l = n.height(),
                        o = $.data(this, d);
                    if (m !== o.w || l !== o.h) { n.trigger(j, [o.w = m, o.h = l]) }
                });
                g()
            }, e[b])
        }
    })(jQuery, this);



    var layout;
    var Ruter = Backbone.Router.extend({
        // Hash maps for routes
        routes: {
            "": "Index",
            "*actions": "defaultRoute"
        },
        Layout: function(amodule, action) {
            var self = this;
            var isLogin = $.cookie('adAuthCookie');
            if (!isLogin) {
                Login = require('login');
                var l = new Login({
                    el: $('body')
                });
                l.render();
            } else {
                var self = this;
                if (!layout) {
                    Layout = require('layout');
                    layout = new Layout({
                        el: $('body')
                    });
                    layout.render(amodule, action, true);
                } else {
                    layout.draw(amodule, action, false);
                }

            }
        },
        Index: function() {
            this.Layout(GodData.Index);
        },
        defaultRoute: function(actions, child) {
            if (actions.indexOf("/") != -1) {
                this.Layout('/bi/biz/' + actions + '.js', actions);
            } else {
                this.Layout('/bi/biz/' + actions + '/' + actions + '.js', actions);
            }
        }
    });
    //自定义验证
    jQuery.validator.methods.compareDate = function(value, element, param) {
        //var startDate = jQuery(param).val() + ":00";补全yyyy-MM-dd HH:mm:ss格式
        //value = value + ":00";

        var startDate = jQuery(param).val();

        var date1 = new Date(Date.parse(startDate.replace("-", "/")));
        var date2 = new Date(Date.parse(value.replace("-", "/")));
        return date1 < date2;
    };
    jQuery.validator.methods.required_super = function(value, element, param) {
        if ($(element).attr('data-id')) {
            return true;
        } else {
            return false;
        }
    };

    $.fn.dataTableExt.oApi.fnStandingRedraw = function(oSettings, isFirstLoad, oTable) {
        if (oSettings.oFeatures.bServerSide === false) {
            var before = oSettings._iDisplayStart;
            oSettings.oApi._fnReDraw(oSettings);
            oSettings._iDisplayStart = before;
            oSettings.oApi._fnCalculateEnd(oSettings);
        }
        if (isFirstLoad) {
            _iDisplayStart = window.localStorage['_iDisplayStart' + oSettings.sAjaxSource]
            _iDisplayLength = window.localStorage['_iDisplayLength' + oSettings.sAjaxSource]
            sSearch = window.localStorage['sSearch' + oSettings.sAjaxSource]
            if (_iDisplayStart != null && _iDisplayLength != null && sSearch != null) {
                if (_iDisplayStart == 0 && _iDisplayStart == 10 && sSearch == '') {
                    return
                }
                if (oTable != null) {
                    oTable.fnClearTable(0)
                }
                oSettings._iDisplayStart = _iDisplayStart
                oSettings._iDisplayLength = _iDisplayLength
                oSettings.oPreviousSearch.sSearch = sSearch
                select = $('.searchDataTableBody').find('select');
                select.find("option[value='" + _iDisplayLength + "']").attr("selected", "selected");
            } else {
                return
            }
        }
        //draw the 'current' page
        oSettings.oApi._fnDraw(oSettings);
    };
    var router = new Ruter;
    Backbone.history.start();
    return {
        Ruter: router
    };
});