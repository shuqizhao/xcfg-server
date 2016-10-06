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
                    'url': '/role/list',
                    'columns': [{
                        'title': 'id',
                        'name': 'Id',
                        'isHide': true
                    }, {
                        'title': '角色名',
                        'name': 'Name'
                    }, {
                        'title': '创建时间',
                        'name': 'CreateTime'
                    }],
                    idName: 'Id',
                    functions: {
                        common: [{
                            text: "添加角色",
                            url: 'auth/roleadd',
                            mode: 'navigate'
                        }]
                    },
                    operations: [{
                        text: '查看',
                        url: 'auth/roleview'
                    }]
                }
            });
            self.searchDataTable.render();
        }
    });
    return mainView;
});