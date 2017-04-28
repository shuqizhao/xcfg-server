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
                    'title': '查看配置文件',
                    'detailTitle': '查看配置文件',
                    'editTitle': '编辑配置文件',
                    'mode': 'detailEdit',
                    'get': {
                        url: '/cfg/get?env=' + self.queryObject.env,
                        params: {
                            id: self.queryObject.id
                        }
                    },
                    'save': '/cfg/update?env=' + self.queryObject.env,
                    'items': [{
                        'name': 'Id',
                        'type': 'hidden'
                    }, {
                        'name': 'CfgName',
                        'title': '文件名',
                        'type': 'text'
                    }, {
                        'name': 'AppName',
                        'title': '应用程序名',
                        'type': 'text'
                    }, {
                        'name': 'CfgFile',
                        'title': '文件内容',
                        'type': 'textxml'
                    }],
                    "afterEditRender": function() {
                        $('#AppName').attr('disabled', true);
                        $('#CfgName').attr('disabled', true);
                    },
                    rules: {
                        CfgName: {
                            required: true
                        },
                        CfgFile: {
                            xml: true
                        }
                    },
                    messages: {
                        CfgName: {
                            required: "配置文件必须填写"
                        }
                    },
                    "validate": function(data) {
                        var result = $.fn.validateXML(data['CfgFile'])
                        self = this
                        if (result.error_code == 0) {
                            return true;
                        } else {
                            $.fn.message({
                                type: 'warning',
                                title: '警告',
                                msg: result.msg
                            });
                            return false;
                        }
                    }
                },

            });
            form.render();
        }
    });
    return mainView;
});