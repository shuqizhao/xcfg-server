<template>
  <list :cfg="cfg"></list>
</template>
<script>
export default {
  data() {
    var self = this;
    return {
      cfg: {
        parentTitle:'配置文件',
        title:'开发配置文件列表',
        url: this.getGlobalData().ApiBaseUrl + "/cfg/list?env="+this.$route.query.env,
        aaSorting: [[6, "desc"]],
        bSort: true,
        columns: [
          {
            title: "Id",
            name: "Id",
            isHide: true
          },
          {
            title: "文件名",
            name: "CfgName",
            isSearch: true
          },
          {
            title: "应用程序名",
            name: "AppName",
            isSearch: true
          },
          {
            title: "版本号",
            name: "MinorVersion"
          },
          {
            title: "类型",
            name: "CfgType",
            isSearch: true,
            type: "combox",
            data: [
              {
                id: -1,
                value: ""
              },
              {
                id: 1,
                value: "标准"
              },
              {
                id: 2,
                value: "Apollo"
              }
            ]
          },
          {
            title: "创建时间",
            name: "CreateTime",
            isSearch: true,
            type: "timer",
            sortable: true
          },
          {
            title: "修改时间",
            name: "UpdateTime",
            isSearch: true,
            type: "timer"
          }
        ],
        fnRowCallback: function(row, data) {
          $("td:eq(2)", row).html(
            '<a href="' +
              self.getGlobalData().ApiBaseUrl +
              "/xcfg/get?id=" +
              data.Id +
              '" target="_blank">' +
              data.CfgName +
              "</a>"
          );
          if (data.CfgType == "" || data.CfgType == "0") {
            $("td:eq(5)", row).html("标准");
          } else {
            $("td:eq(5)", row).html("Apollo");
          }
        },
        idName: "Id",
        functions: {
          common: [
            {
              text: "新建标准配置文件",
              url: "/xcfg/xcfgadd?env="+this.$route.query.env,
              mode: "navigate"
            },
            {
              text: "新建Apollo模板",
              url: "/xcfg/xcfgapolloadd?env="+this.$route.query.env,
              mode: "navigate"
            }
          ]
        },
        operations: [
          {
            text: "查看",
            url: "/xcfg/xcfgview?env="+this.$route.query.env
          },
          {
            text: "历史",
            url: "/xcfg/xcfghistoryview?env="+this.$route.query.env
          }
        ]
      }
    };
  }
};
</script>