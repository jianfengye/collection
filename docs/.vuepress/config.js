module.exports = {
  title: "collection库", // 设置网站标题
  description: "一个让业务开发效率提高10倍的golang库", //描述
  dest: "../../dist/", // 设置输出目录
  port: 2333, //端口
  head: [["link", { rel: "icon", href: "/assets/img/head.png" }]],
  themeConfig: {
    // 添加导航栏
    nav: [
      { text: "主页", link: "/" }, // 导航条
      { text: "指南", link: "/guide/install" },
      { text: "手册", link: "/method/append" },
      {
        text: "github",
        // 这里是下拉列表展现形式。
        items: [
          {
            text: "collection",
            link: "https://github.com/jianfengye/collection",
          },
        ],
      },
    ],
    // 为以下路由添加侧边栏
    sidebar: {
      "/method/": [
        {
          title: "手册",
          collapsable: false,
          children: [
            "append",
            "avg",
            "contains",
            "containscount",
            "copy",
            "dd",
            "diff",
            "each",
            "every",
            "filter",
            "first",
            "forpage",
            "myindex",
            "isempty",
            "isnotempty",
            "join",
            "last",
            "map",
            "max",
            "median",
            "merge",
            "min",
            "mod",
            "newempty",
            "nth",
            "pad",
            "pluck",
            "pop",
            "prepend",
            "push",
            "random",
            "reduce",
            "reject",
            "reverse",
            "search",
            "shuffle",
            "slice",
            "sort",
            "sortby",
            "sortbydesc",
            "sortdesc",
            "sum",
            "tofloat32s",
            "tofloat64s",
            "toint64s",
            "tointerfaces",
            "toints",
            "tomixs",
            "unique",
            "toobjects",
            "split",
            "groupby"
          ],
        },
      ],
      "/guide/": [
        {
          title: "指南",
          collapsable: true,
          children: [
            "install",
            "introduce",
            "benchmark",
            "v1.3.1",
            "v1.4.0"
          ],
        },
      ],
    },
  },
};
