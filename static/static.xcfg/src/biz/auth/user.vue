<template>
  <list :cfg="cfg"></list>
</template>
<script>
export default {
  data() {
    var self = this;
    return {
      cfg: {
        title: "用户管理列表",
        parentTitle:'权限管理',
        url: this.getGlobalData().ApiBaseUrl +"/user/list",
        columns: [
          {
            title: "用户名",
            name: "Name",
            isSearch: true
          },
          {
            title: "是否启用",
            name: "IsEnabled",
            isSearch: true,
            type: "combox",
            data: [
              {
                id: "2",
                value: "不限"
              },
              {
                id: "1",
                value: "是"
              },
              {
                id: "0",
                value: "否"
              }
            ]
          },
          {
            title: "创建时间",
            name: "CreateTime",
            isSearch: true,
            type: "timer"
          },
          {
            title: "最后登录时间",
            name: "LastLoginTime"
          }
        ],
        fnRowCallback: function(row, data) {
          if (data.IsEnabled) {
            $("td:eq(3)", row).html('<i class="fa fa-fw fa-check-circle"></i>');
          } else {
            $("td:eq(3)", row).html('<i class="el-icon-close"></i>');
          }
        },
        idName: "Id",
        functions: {
          more: [
            {
              text: "停用",
              url: this.getGlobalData().ApiBaseUrl +"/user/disable"
            },
            {
              text: "启用",
              url: this.getGlobalData().ApiBaseUrl +"/user/enable"
            }
          ],
          common: [
            {
              text: "添加用户",
              url: "/auth/useradd",
              mode: "navigate"
            }
          ]
        },
        operations: [
          {
            text: "查看",
            url: "/auth/userview"
          }
        ]
      }
    };
  }
};
</script>