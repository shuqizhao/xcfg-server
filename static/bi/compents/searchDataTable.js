define(['backbone', 'text!searchDataTable.html', 'godModel', 'dataTable', 'modalForm'], function(Backbone, jst, GodModel, DataTable, ModalForm) {
    var mainView = Backbone.View.extend({
        template: _.template(jst),
        initialize: function(options) {
            this.godModel = new GodModel();
            this.options = options;
        },
        events: {
            'click .btn-searchDataTable': 'doSearch',
            'click .searchDataTableCheckAll': 'searchDataTableCheckAll',
            'click .searchDataTableMoreOp': 'searchDataTableMoreOp',
            'click .searchDataTableDetail': 'searchDataTableDetail',
            'click .dropdown-toggle': 'dropdowntoggle',
            'click .caret': 'caret'
        },
        render: function() {
            var self = this;
            var cfg = self.options.cfg;
            var searchColumns = [];
            var aoColumns = [];
            if (cfg.functions) {
                aoColumns.push({
                    "sTitle": "<center><input type=\"checkbox\" class=\"searchDataTableCheckAll\" /></center>",
                    "mDataProp": [0],
                    "sName": "",
                    "bSearchable": false,
                    "bSortable": false,
                    "fnRender": function(oObj) {
                            var str = '<center><input type=\"checkbox\" value=\"' + oObj.aData[cfg.idName] + '\" /></center>';
                            return str;
                        } //自定义列的样式
                });
            }
            if (cfg.operations) {
                aoColumns.push({
                    "sTitle": "操作",
                    "mDataProp": [1],
                    "sName": "",
                    "bSearchable": false,
                    "bSortable": false,
                    "fnRender": function(oObj) {
                            var str = '';
                            for (var i = 0; i < cfg.operations.length; i++) {
                                var op = cfg.operations[i];
                                if (i != 0) {
                                    str += ' | ';
                                }
                                var idName = op.idName ? op.idName : cfg.idName;
                                var presuff = "";
                                if (op.url.indexOf('frame?url=') >= 0 || op.url.indexOf('address?url=') >= 0) {
                                    presuff = "?";
                                } else {
                                    presuff = op.url.indexOf('?') > 0 ? '&' : '?';
                                }
                                var paramUrl = presuff + 'id=' + oObj.aData[idName];
                                if (cfg.urlParams) {
                                    for (var j = 0; j < cfg.urlParams.length; j++) {
                                        paramUrl += "&" + cfg.urlParams[j] + "=" + oObj.aData[cfg.urlParams[j]];
                                    }
                                }
                                if (!op.target) {
                                    op.target = "_self";
                                }
                                if (op.url.indexOf('address?url=') >= 0) {
                                    str += '<a href=\"/' + op.url.replace('address?url=', '') + paramUrl + '\" target=\"' + op.target + '\">' + op.text + '</a>';
                                } else {
                                    str += '<a href=\"#' + op.url + paramUrl + '\" target=\"' + op.target + '\">' + op.text + '</a>';
                                }
                            };
                            return str;

                        } //自定义列的样式
                });
            }
            for (var i = 0; i < cfg.columns.length; i++) {
                var column = cfg.columns[i];
                if (column.isSearch) {
                    searchColumns.push(column);
                }
                if (!column.isHide) {
                    aoColumns.push({
                        "sTitle": column.title,
                        "mDataProp": column.name,
                        "sName": column.name,
                        "sWidth": column.width
                    });
                }
            };

            var data = {
                'columns': searchColumns,
                'title': cfg.title
            };
            self.$el.html(self.template(data));

            var searchDataTableTop = self.$el.find('.searchDataTableTop');
            if (searchColumns.length == 0) {
                searchDataTableTop.hide();
            }

            var span10Height = searchDataTableTop.find('.span10').height();
            searchDataTableTop.find('.span2').height(span10Height);

            var searchButton = searchDataTableTop.find('.btn-searchDataTable');

            var searchButtonHeight = searchButton.outerHeight();
            searchButton.css('margin-top', (span10Height - searchButtonHeight) / 2 + 'px');


            var dataTableCfg = {
                "sAjaxSource": cfg.url,
                "aoColumns": aoColumns,
                "aoColumnDefs": [{
                    sDefaultContent: '',
                    aTargets: ['_all']
                }],
                "fnDrawCallback": function(oSettings) {
                    //console.log(oSettings)
                    //var iCurrentPage = Math.ceil(oSettings._iDisplayStart / oSettings._iDisplayLength) + 1;
                    //console.log(iCurrentPage)
                    window.localStorage['_iDisplayStart' + oSettings.sAjaxSource] = oSettings._iDisplayStart
                    window.localStorage['_iDisplayLength' + oSettings.sAjaxSource] = oSettings._iDisplayLength
                    window.localStorage['sSearch' + oSettings.sAjaxSource] = oSettings.oPreviousSearch.sSearch
                    self.$el.find('.searchDataTableCheckAll').attr("checked", false)
                }
            };
            var lastCfg = $.extend(true, dataTableCfg, self.options.cfg);
            var dataTable = new DataTable({
                el: self.$el.find('.searchDataTableBody'),
                cfg: lastCfg
            });

            dataTable.render();
            self.$el.find(".dataTables_filter").hide();
            self.dataTable = dataTable;
            var buttons = "";
            //功能按钮处理
            var functions = self.options.cfg.functions;
            if (functions) {
                if (functions && functions.common) {
                    for (var i = 0; i < functions.common.length; i++) {
                        var button = functions.common[i];
                        buttons += "<button id='" + button.text + "' class='btn btn-success searchDataTableMoreOp' url='" + button.url + "' mode='" + button.mode + "' style='margin-left:10px;' limitSelected='" + button.limitSelected + "'>" + button.text + "</button>";
                    };
                }
                if (functions.more) {
                    buttons += '<div class="btn-group" style="margin-left:10px;" >'
                    buttons += '<button type="button" class="btn btn-success">更多操作</button>'
                    buttons += '<button type="button" class="btn btn-success dropdown-toggle" data-toggle="dropdown">'
                    buttons += '<span class="caret"></span>'
                    buttons += '<span class="sr-only"></span>'
                    buttons += '</button>'
                    buttons += '<ul class="dropdown-menu dropdown-ul-menu" role="menu">'
                    if (functions && functions.more) {
                        for (var i = 0; i < functions.more.length; i++) {
                            var button = functions.more[i];
                            buttons += '<li><a class="searchDataTableMoreOp" id="' + button.text + '" url="' + button.url + '" mode="' + button.mode + '" limitSelected=' + button.limitSelected + '>' + button.text + '</a></li>';
                            buttons += '<li class="divider"></li>';
                        };
                    }
                    buttons += '</ul>';
                    buttons += '</div>';
                }
            }
            self.$el.find(".dataTables_function").html(buttons);
            self.$el.resize(function() {
                var searchDataTableTop = self.$el.find('.searchDataTableTop');

                var span10Height = searchDataTableTop.find('.span10').height();
                searchDataTableTop.find('.span2').height(span10Height);

                var searchButton = searchDataTableTop.find('.btn-searchDataTable');

                var searchButtonHeight = searchButton.outerHeight();
                searchButton.css('margin-top', (span10Height - searchButtonHeight) / 2 + 'px');
            });
        },
        doSearch: function() {
            var self = this;
            var data = {};
            self.$el.find(".searchDataTable").find(".form-control").each(function(index) {
                var item = $(this);
                if (item.attr('controltype') == "suggest") {
                    data[this.id] = item.attr('data-id');
                } else {
                    data[this.id] = item.val();
                }
            });
            this.dataTable.sourceDataTable.fnFilter(JSON.stringify(data));
        },
        searchDataTableCheckAll: function(e) {
            if ($(e.target).attr("checked")) {
                this.$el.find(".searchDataTable").find(":checkbox").attr("checked", true);
            } else {
                this.$el.find(".searchDataTable").find(":checkbox").attr("checked", false);
            }

        },
        searchDataTableMoreOp: function(e) {
            var self = this;
            var mode = $(e.target).attr("mode");
            var url = $(e.target).attr("url");
            var limitSelected = $(e.target).attr("limitSelected");
            if (mode == 'navigate') {
                Backbone.history.navigate(url, {
                    trigger: true,
                    replace: false
                });
                return;
            } else if (mode == 'modal') {
                var modalForm = new ModalForm({
                    el: self.$el,
                    url: url
                });
                modalForm.render();
                Backbone.history.navigate(url);
                return;
            }
            var self = this;
            var tipMsg = $(e.target).text();
            var checks = self.$el.find(".searchDataTable").find(":checkbox[checked]");
            if (limitSelected) {
                if (limitSelected <= checks.length-1) {
                    $.fn.message({
                        msg: "只能选择" + limitSelected + "个对象！",
                        type: "warning"
                    });
                    return;
                }
            }
            if (mode != 'skipcheck' && checks.length == 0) {

                $.fn.message({
                    msg: "请选择" + tipMsg + "对象！",
                    type: "warning"
                });
            } else {
                $.messager.model = {
                    ok: {
                        text: "确认",
                        classed: 'btn-primary'
                    },
                    cancel: {
                        text: "取消",
                        classed: 'btn-danger'
                    }
                };
                $.messager.confirm("提示", "<h3>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;确定要" + tipMsg + "吗？</h3>", function() {
                    var data = [];
                    checks.each(function() {
                        var idValue = $(this).val();
                        if (idValue != 'on') {
                            data.push($(this).val());
                        }
                    });
                    if (mode == 'download') {
                        if (limitSelected) {
                            if (limitSelected < data.length) {
                                $.fn.message({
                                    msg: "只能下载" + limitSelected + "个对象！",
                                    type: "warning"
                                });
                                return;
                            }
                        }
                        var form = $("<form>");
                        form.attr("style", "display:none");
                        form.attr("target", "");
                        form.attr("method", "post");
                        form.attr("action", url);
                        for (var i = 0; i < data.length; i++) {
                            var input1 = $("<input>");
                            input1.attr("type", "hidden");
                            input1.attr("name", "ids");
                            input1.attr("value", data[i]);
                            $("body").append(form);
                            form.append(input1);
                        }
                        form.submit();
                    } else {
                        self.godModel.url = url;
                        self.godModel.set({
                            "ids": data
                        });
                        self.godModel.save({}, {
                            success: function(model, response) {
                                if (response.code == 200) {
                                    self.dataTable.sourceDataTable.fnStandingRedraw(false);
                                    $.fn.message({
                                        msg: tipMsg + "成功！"
                                    });
                                } else {
                                    $.fn.message({
                                        msg: response.message + "！",
                                        type: "error"
                                    });
                                }
                            },
                            error: function(err) {
                                $.fn.message({
                                    msg: tipMsg + "失败！",
                                    type: "error"
                                });
                            }
                        });
                    }
                });
            }
        },
        reLoad: function() {
            this.dataTable.sourceDataTable.fnFilter("");
        },
        searchDataTableDetail: function() {
            //alert(10)
        },
        dropdowntoggle: function(e) {
            $(e.target).next().show();
        },
        caret: function(e) {
            $(e.target).parent().next().show();
        }
    });
    return mainView;
});