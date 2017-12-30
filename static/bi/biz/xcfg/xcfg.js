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
                    'url': '/cfg/list?env=' + self.queryObject.env,
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
                    },{
                        'title': '类型',
                        'name': 'CfgType'
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
                        $('td:eq(2)', row).html('<a href="/xcfg/get?id=' + data.Id + '" target="_blank">' + data.CfgName + '</a>');
                        if(data.CfgType=="" || data.CfgType=="0"){
                            $('td:eq(5)', row).html('标准');
                        }else{
                            $('td:eq(5)', row).html('Apollo');
                        }
                    },
                    idName: 'Id',
                    functions: {
                        common: [{
                            text: "新建标准配置文件",
                            url: 'xcfg/xcfgadd?env=' + self.queryObject.env,
                            mode: 'navigate'
                        },{
                            text: "新建Apollo模板",
                            url: 'xcfg/xcfgapolloadd?env=' + self.queryObject.env,
                            mode: 'navigate'
                        }]
                    },
                    operations: [{
                        text: '查看',
                        url: 'xcfg/xcfgview?env=' + self.queryObject.env
                    },{
                        text: '历史',
                        url: 'xcfg/xcfghistoryview?env=' + self.queryObject.env
                    }],
                }
            });
            self.searchDataTable.render();
        }
    });
    return mainView;
});