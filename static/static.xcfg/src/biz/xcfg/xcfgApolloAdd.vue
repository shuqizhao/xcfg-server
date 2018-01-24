<template>
 <mform :cfg="cfg"></mform>
</template>
<script>
export default {
  data() {
    var self = this;
    return {
      cfg: {
        title: "新建Apollo模板",
        mode: "create",
        save:
          this.getGlobalData().ApiBaseUrl +
          "/cfg/add?env=" +
          +this.$route.query.env +
          "&type=1",
        items: [
          {
            name: "Id",
            type: "hidden"
          },
          {
            name: "CfgName",
            title: "文件名",
            type: "text"
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
        onSuccess: function(mode, response) {
          if (response.code == 200) {
            history.back();
          } else {
            self.$message({
              message: "添加配置文件失败！",
              type: "error"
            });
          }
        },
        validate: function(data, saveData) {
          var result = $.fn.validateXML(data["CfgFile"]);
          if (result.error_code == 0) {
            $.ajax({
              type: "POST",
              xhrFields: {
                withCredentials: true
              },
              url:
                self.getGlobalData().ApiBaseUrl +
                "/cfg/exists?env=" +
                self.$route.query.env,
              data: data,
              success: function(response) {
                if (response.code == 200 && response.data == 0) {
                  saveData(data);
                } else {
                  self.$message({
                    type: "warning",
                    messages: "配置文件已经存在!"
                  });
                }
              }
            });
          } else {
            self.$message({
              type: "warning",
              message: result.msg
            });
          }
          return false;
        }
      }
    };
  }
};
</script>