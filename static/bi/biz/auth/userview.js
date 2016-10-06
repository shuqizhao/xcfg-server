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
                    'title': '查看用户',
                    'detailTitle': '查看用户',
                    'editTitle': '编辑用户的角色',
                    'mode': 'detailEdit',
                    'get': {
                        url: '/user/get',
                        params: {
                            id: self.queryObject.id
                        }
                    },
                    'save': '/user/update',
                    'items': [{
                        'name': 'Id',
                        'type': 'hidden'
                    }, {
                        'name': 'Name',
                        'title': '用户名',
                        'type': 'text'
                    }, {
                        'name': 'Roles',
                        'title': '角色',
                        'type': 'select2select',
                        'url': "/user/roles?id=" + self.queryObject.id
                    }, {
                        'name': 'IsEnabled',
                        'title': '是否启用',
                        'type': 'yesno'
                    }],
                    "afterEditRender": function() {
                        self.$el.find('#Name').attr('disabled', true);
                        self.$el.find('#IsEnabled').attr('disabled', true);
                    },
                },

            });
            form.render();
        }
    });
    return mainView;
});