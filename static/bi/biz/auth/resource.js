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
                    'url': '/menu/list',
                    'columns': [{
                        'title': 'id',
                        'name': 'Id',
                        'isHide': true
                    }, {
                        'title': '资源名',
                        'name': 'Name'
                    }, {
                        'title': 'URL',
                        'name': 'Url'
                    }],
                    idName: 'Id'
                }
            });
            self.searchDataTable.render();
        }
    });
    return mainView;
});