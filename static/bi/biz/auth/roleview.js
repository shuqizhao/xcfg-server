define(['backbone', 'form'], function(Backbone, Form) {
    var mainView = Backbone.View.extend({
        initialize: function(options) {
            this.options = options;
            this.queryObject = $.fn.getQueryObject();
        },
        render: function() {
            var self = this;
            var form = new Form({
                el: self.el,
                cfg: {
                    'title': '查看角色',
                    'detailTitle': '查看角色',
                    'editTitle': '编辑角色的资源',
                    'mode': 'detailEdit',
                    'get': {
                        url: '/role/get',
                        params: {
                            id: self.queryObject.id
                        }
                    },
                    'save': '/role/update',
                    'items': [{
                        'name': 'Id',
                        'type': 'hidden'
                    }, {
                        'name': 'Name',
                        'title': '角色名',
                        'type': 'text'
                    }, {
                        'name': 'Menus',
                        'title': '资源',
                        'type': 'select2select',
                        'url': "/role/resources?id=" + self.queryObject.id
                    }],
                    "afterEditRender": function() {
                        self.$el.find('#Name').attr('disabled', true);
                    },
                },

            });
            form.render();
        }
    });
    return mainView;
});