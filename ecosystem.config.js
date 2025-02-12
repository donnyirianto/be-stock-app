module.exports = {
    apps: [
      {
        name: "stock-app",
        script: "/root/be-stock-app/stock-app -config=/root/be-stock-app",
        exec_mode: "fork",
        autorestart: true,
        watch: false,
        max_memory_restart: "900M",
        env: {
          APP_ENV: "production",
        },
        log_date_format: "YYYY-MM-DD HH:mm:ss:SSS Z",
        merge_logs: true,
      },
    ],
  };