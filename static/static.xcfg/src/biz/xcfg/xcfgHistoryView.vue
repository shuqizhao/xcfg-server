<template>
  <list :cfg="cfg"></list>
</template>
<script>
export default {
  data() {
    var self = this;
    return {
      cfg: {
        url:
          self.getGlobalData().ApiBaseUrl +
          "/cfghistory/list?id=" +
          self.$route.query.id,
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
            title: "创建时间",
            name: "CreateTime",
            isSearch: true,
            type: "timer"
          },
          {
            title: "修改时间",
            name: "UpdateTime",
            isSearch: true,
            type: "timer"
          }
        ],
        fnRowCallback: function(row, data) {
          if (data.CfgType == "1") {
            $("td:eq(1)", row).html(
              '<a href="' +
                self.getGlobalData().ApiBaseUrl +
                "/cfghistory/get?id=" +
                data.Id +
                '" target="_blank">' +
                data.CfgName +
                '</a>&nbsp;&nbsp;<a href="' +
                self.getGlobalData().ApiBaseUrl +
                "/cfghistory/getTemplate?id=" +
                data.Id +
                '" target="_blank">模板</a>'
            );
          } else {
            $("td:eq(1)", row).html(
              '<a href="' +
                self.getGlobalData().ApiBaseUrl +
                "/cfghistory/get?id=" +
                data.Id +
                '" target="_blank">' +
                data.CfgName +
                "</a>"
            );
          }
        },
        idName: "Id",
        functions: {
          common: [
            {
              text: "回滚",
              url: self.getGlobalData().ApiBaseUrl + "/cfghistory/rollback",
              limitSelected: 1
            }
          ]
        }
      }
    };
  }
};
</script>