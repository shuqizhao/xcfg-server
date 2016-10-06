define(['backbone', 'godModel', 'jDataTable', 'bDataTable'], function(Backbone, GodModel) {
	var mainView = Backbone.View.extend({
		initialize: function(options) {
			this.options = options;
		},
		render: function() {
			var self = this;
			var tableHtml = '<div class="container-fluid"><div class="row-fluid"><table class="table table-striped table-bordered table-hover datatable"></table></div></div>';
			self.$el.html(tableHtml);
			var cfg = {
				"oLanguage": {
					"sProcessing": "处理中...",
					"sLengthMenu": "每页 _MENU_ 条记录",
					"sZeroRecords": "没有匹配的记录",
					"sInfo": "显示第 _START_ 至 _END_ 项记录，共 _TOTAL_ 项",
					"sInfoEmpty": "显示第 0 至 0 项记录，共 0 项",
					"sInfoFiltered": "(由 _MAX_ 项记录过滤)",
					"sInfoPostFix": "",
					"sSearch": "过滤:",
					"sUrl": "",
					"oPaginate": {
						"sFirst": "首页",
						"sPrevious": "上页",
						"sNext": "下页",
						"sLast": "末页"
					}
				},
				"bRetrieve": true,
				'bFilter': true,
				"bPaginate": true,
				"bServerSide": true,
				"bLengthChange": true,
				"sServerMethod": "POST",
				"sDom": '<"H"<"dataTables_function"/>rp>t<"F"lip>'
			};
			var lastCfg = $.extend(true, cfg, self.options.cfg);
			self.sourceDataTable = $('.datatable').dataTable(lastCfg);
		}
	});
	return mainView;
});