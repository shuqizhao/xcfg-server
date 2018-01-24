<template>
 <mform :cfg="cfg"></mform>
</template>
<script>
export default {
  data() {
    var self = this;
    return {
      cfg: {
        title: "查看用户",
        detailTitle: "查看用户",
        editTitle: "编辑用户的角色",
        mode: "detailEdit",
        get: {
          url: this.getGlobalData().ApiBaseUrl + "/user/get",
          params: {
            id: this.$route.query.id
          }
        },
        save: this.getGlobalData().ApiBaseUrl + "/user/update",
        items: [
          {
            name: "Id",
            type: "hidden"
          },
          {
            name: "Name",
            title: "用户名",
            type: "text"
          },
          {
            name: "Roles",
            title: "角色",
            type: "select2select",
            url:
              this.getGlobalData().ApiBaseUrl +
              "/user/roles?id=" +
              this.$route.query.id
          },
          {
            name: "IsEnabled",
            title: "是否启用",
            type: "yesno"
          }
        ],
        afterEditRender: function(mode,data) {
          $(self.$el).find("#Name").attr("disabled", true);
          $(self.$el).find("#IsEnabled").attr("disabled", true);
           $(self.$el).find("#IsEnabled input").attr("disabled", true);
        }
      }
    };
  }
};
</script>