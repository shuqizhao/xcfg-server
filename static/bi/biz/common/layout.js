define(['backbone', 'text!layout.html', 'tree', 'godModel', 'toastmessage', 'validate', 'message', 'jBootstrap', 'jDataTable', 'select2select'].concat(GodData.RegistBiz), function(Backbone, jst, Tree, GodModel) {
    var mainView = Backbone.View.extend({
        template: _.template(jst),
        initialize: function(options) {
            this.options = options;
            this.godModel = new GodModel();
        },
        events: {
            'click #logout': 'logout'
        },
        render: function(amodule, action, first) {
            var self = this;
            document.title = "远程配置文件管理系统";
            var data = { user: $.cookie('loginUser') };
            self.$el.html(this.template(data));

            self.godModel = new GodModel();
            self.godModel.url = '/auth/menus';
            self.godModel.fetch({
                success: function(model, text) {
                    result = model.toJSON();
                    self.dataofTree = [];
                    parentMenus = []
                    for (i in result.data) {
                        if (result.data[i].ParentId == 0) {
                            parentMenus.push(result.data[i])
                        }
                    }
                    for (i in parentMenus) {
                        pm = parentMenus[i]
                        cms = []
                        for (j in result.data) {
                            cm = result.data[j]
                            if (cm.ParentId == pm.Id) {
                                cms.push({ "Name": cm.Name, "Url": cm.Url })
                            }
                        }
                        self.dataofTree.push({
                            "Name": pm.Name,
                            "HasChild": true,
                            "Childs": cms
                        });
                    }

                    self.tree = new Tree({
                        el: $('.left'),
                        mode: 'cfg',
                        data: {
                            "Name": "所有功能",
                            "HasChild": true,
                            "Childs": self.dataofTree
                        }
                    });
                    self.tree.render();
                    self.draw(amodule, action, first)
                }
            });

            self.$el.find(".dropdown-toggle").click(function(e) {
                $(e.target).next().show();
            });
            self.$el.find(".dropdown-menu").mouseleave(function(e) {
                $(this).hide();
            });
            $(document).bind("click", function(e) {
                self.$el.find(".dropdown-ul-menu").hide();
            });
            self.$el.find(".center").resize(function() {
                var item = self.$el.find("input[controltype='suggest']");
                if (item.attr("controltype") == "suggest") {
                    item.next().find(".dropdown-menu").css("left", item.position().left);
                    item.next().find(".dropdown-menu").css("top", item.position().top + 30);
                    item.next().find(".dropdown-menu").css("min-width", "280px");
                }
            });
        },
        draw: function(amodule, action, first) {
            $('.showbox').remove();
            $('.center').html('<div class="showbox"></div>');
            this.tree.selectNode(action, first);
            var Module = require("/static" + amodule);
            var m = new Module({
                el: $('.showbox')
            });
            m.render();
        },
        logout: function() {
            $.cookie('adAuthCookie', "false", { expires: -1, path: "/" });
            $.cookie('loginUser', "false", { expires: -1, path: "/" });
            location.reload();
        }
    });
    return mainView;
});