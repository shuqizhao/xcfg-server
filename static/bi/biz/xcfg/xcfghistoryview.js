define(['backbone', 'searchDataTable'], function(Backbone, SearchDataTable) {
    var mainView = Backbone.View.extend({
        initialize: function(options) {
            this.options = options;
            this.queryObject = $.fn.getQueryObject();
        },
        render: function() {
            var self = this;
            self.searchDataTable = new SearchDataTable({
                el: self.el,
                cfg: {
                    'url': '/cfghistroy/list?id=' + self.queryObject.id,
                    'columns': [{
                        'title': 'Id',
                        'name': 'Id',
                        'isHide': true
                    }, {
                        'title': '文件名',
                        'name': 'CfgName',
                        'isSearch': true
                    }, {
                        'title': '应用程序名',
                        'name': 'AppName',
                        'isSearch': true
                    }, {
                        'title': '版本号',
                        'name': 'MinorVersion'
                    }, {
                        'title': '创建时间',
                        'name': 'CreateTime',
                        'isSearch': true,
                        'type': 'timer'
                    }, {
                        'title': '修改时间',
                        'name': 'UpdateTime',
                        'isSearch': true,
                        'type': 'timer'
                    }],
                    "fnRowCallback": function(row, data) {
                        $('td:eq(1)', row).html('<a href="/cfghistory/get?id=' + data.Id + '" target="_blank">' + data.CfgName + '</a>');
                    },
                    idName: 'Id',
                    functions: {
                        common: [{
                            text: "回滚",
                            url: 'cfghistroy/rollback',
                            limitSelected:1
                        }]
                    },
                }
            });
            self.searchDataTable.render();
        }
    });
    return mainView;
});