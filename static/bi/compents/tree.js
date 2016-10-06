define(['backbone', 'godModel'], function(Backbone, GodModel) {
	var mainView = Backbone.View.extend({
		initialize: function(options) {
			this.treeModel = new GodModel();
			this.options = options;
		},
		events: {
			'click .treeNode': 'treeNodeClick',
			'click .btn-hide': 'btnHide',
		},
		registEvent: function() {
			$('.tree li:has(ul)').addClass('parent_li').find(' > span').attr('title', '关闭节点');
			$('.tree li.parent_li > span').on('click', function(e) {
				var children = $(this).parent('li.parent_li').find(' > ul > li');
				if (children.is(":visible")) {
					children.hide('fast');
					$(this).attr('title', '展开节点').find(' > i').addClass('icon-plus-sign').removeClass('icon-minus-sign');
				} else {
					children.show('fast');
					$(this).attr('title', '关闭节点').find(' > i').addClass('icon-minus-sign').removeClass('icon-plus-sign');
				}
				e.stopPropagation();
			});
		},
		loadData: function() {
			var self = this;
			if (self.options.mode == 'url') {
				self.treeModel.url = self.options.url;
				self.treeModel.fetch({
					success: function(model, text) {
						var data = model.toJSON();
						var html = self.buildTree(self.options.cfg, data);
						self.$el.html(html);
						self.registEvent();
					},
					error: function() {}
				});
			} else {
				var html = self.buildTree({}, self.options.data);
				self.$el.html(html);
				self.registEvent();
			}

		},
		buildTree: function(cfg, data) {
			var html = '<div class="tree well"><ul style="margin:0px;">';
			html += this.buildTreeNode(cfg, data);
			html += '</ul>';
			html += '<button class="btn btn-primary btn-hide" style="width: 100%;">隐藏</button>';
			html += '</div>';
			return html;
		},
		buildTreeNode: function(cfg, data) {
			var html = '<li>';
			var self = this;
			var rootName = 'Childs';

			if (data.HasChild) {
				html += '<span><i class="icon-minus-sign"></i>' + data.Name + '</span>';
				html += '<ul>';
				var childs = data[rootName];
				for (var child in childs) {
					html += self.buildTreeNode(cfg, childs[child]);
				}
				html += '</ul>';
			} else {
				html += '<span><i class="icon-leaf"></i><a class="treeNode" href="#' + data.Url + '">' + data.Name + '</a></span>';
			}
			html += '</li>';
			return html;
		},
		render: function() {
			this.loadData();
		},
		treeNodeClick: function(e) {
			$('.treeNode').each(function() {
				$(this).parent().css('background-color', '#fbfbfb');
			});
			$(e.target).parent().css('background-color', '#94a0b4');
		},
		selectNode: function(action, first) {
			if (first) {
				$('.treeNode').each(function() {
					$(this).parent().css('background-color', '#fbfbfb');
				});
				$('a[href="#' + action + '"]').parent().css('background-color', '#94a0b4');
			}
		},
		btnHide: function(e) {
			var target = $(e.target);
			var text = target.text();
			if (text == '隐藏') {
				this.treeDo(true);
				target.text("显示");
				$('i[class="icon-minus-sign"]').removeClass("icon-minus-sign").addClass("icon-plus-sign");
				var left = target.parent().parent();
				left.removeClass('span3').addClass('span1');
				left.next().removeClass('span9').addClass('span11');
			} else {
				this.treeDo(false);
				target.text("隐藏");
				$('i[class="icon-plus-sign"]').removeClass("icon-plus-sign").addClass("icon-minus-sign");
				var left = target.parent().parent();
				left.removeClass('span1').addClass('span3');
				left.next().removeClass('span11').addClass('span9');
			}
		},
		treeDo: function(isHiden) {
			$('.tree li.parent_li > span').each(function() {
				var children = $(this).parent('li.parent_li').find(' > ul > li');
				if (isHiden) {
					children.hide('fast');
					$(this).attr('title', '展开节点').find(' > i').addClass('icon-plus-sign').removeClass('icon-minus-sign');
				} else {
					children.show('fast');
					$(this).attr('title', '关闭节点').find(' > i').addClass('icon-minus-sign').removeClass('icon-plus-sign');
				}
			});
		}
	});
	return mainView;
});