define(['backbone', 'searchDataTable'], function(Backbone, SearchDataTable) {
    var mainView = Backbone.View.extend({
        initialize: function(options) {
            this.options = options;
        },
        render: function() {
            var self = this;
            self.searchDataTable = new SearchDataTable({
                el: self.el,
                cfg: {
                    'url': '/user/list',
                    'columns': [{
                        'title': '用户名',
                        'name': 'Name',
                        'isSearch': true
                    }, {
                        'title': '是否启用',
                        'name': 'IsEnabled',
                        'isSearch': true,
                        'type': 'combox',
                        data: [{
                            id: '2',
                            value: '不限'
                        }, {
                            id: '1',
                            value: '是'
                        }, {
                            id: '0',
                            value: '否'
                        }]
                    }, {
                        'title': '创建时间',
                        'name': 'CreateTime',
                        'isSearch': true,
                        'type': 'timer'
                    }, {
                        'title': '最后登录时间',
                        'name': 'LastLoginTime'
                    }],
                    "fnRowCallback": function(row, data) {
                        if (data.IsEnabled) {
                            $('td:eq(3)', row).html('<b class="icon-ok"></b>');
                        } else {
                            $('td:eq(3)', row).html('<b class="icon-remove"></b>');
                        }
                    },
                    idName: 'Id',
                    functions: {
                        more: [{
                            text: "停用",
                            url: 'user/disable'
                        }, {
                            text: "启用",
                            url: 'user/enable'
                        }],
                        common: [{
                            text: "添加用户",
                            url: 'auth/useradd',
                            mode: 'navigate'
                        }]
                    },
                    operations: [{
                        text: '查看',
                        url: 'auth/userview'
                    }]
                }
            });
            self.searchDataTable.render();
        }
    });
    return mainView;
});