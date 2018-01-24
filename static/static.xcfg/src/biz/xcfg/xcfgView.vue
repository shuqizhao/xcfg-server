<template>
    <mform :cfg="cfg"></mform>
</template>
<script>
export default {
  data() {
    return {
      cfg: {
        title: "查看配置文件",
        detailTitle: "查看配置文件",
        editTitle: "编辑配置文件",
        mode: "detailEdit",
        get: {
          url:  this.getGlobalData().ApiBaseUrl +"/cfg/get?env="+this.$route.query.env,
          params: {
            id: this.$route.query.id
          }
        },
        save: this.getGlobalData().ApiBaseUrl +"/cfg/update?env=dev",
        items: [
          {
            name: "Id",
            type: "hidden"
          },
          {
            name: "CfgName",
            title: "文件名",
            type: "text",
            helpblock:'文件必填*,大于5小于6'
          },
          {
            name: "AppName",
            title: "应用程序名",
            type: "text"
          },
          {
            name: "CfgFile",
            title: "文件内容",
            type: "textxml"
          },
          {
            name: "ApolloTemplate",
            title: "模板内容",
            type: "textnginx"
          }
        ],
        afterEditRender: function(model, data) {
          if (model == "edit") {
            $("#AppName").attr("disabled", true);
            $("#CfgName").attr("disabled", true);
            if (data.CfgType != "1") {
              $("#ApolloTemplate")
                .parent()
                .parent()
                .hide();
            }
          } else {
            if (data.CfgType != "1") {
              $('.form-group :contains("模板内容")')
                .parent()
                .hide();
            }
          }
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
        validate: function(data) {
          var result = self.validateXML(data["CfgFile"]);
          if (result.error_code == 0) {
            return true;
          } else {
            self.$message({
              message: result.msg,
              type: 'warning'
            });
            return false;
          }
        }
      }
    };
  }
};
</script>

