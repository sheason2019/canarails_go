import { Typography } from '@mui/material';
import DependenciesGroup from './components/dependencies-group';

export default function Dependencies() {
  return (
    <>
      <Typography sx={{ mt: 3, mb: 1 }} variant="h5">
        依赖信息
      </Typography>
      <Typography>
        下面列举了 Canarails 主要的依赖项，以及与之相关的网页链接。
      </Typography>
      <DependenciesGroup
        label="Frontend"
        items={[
          {
            label: 'React - 视图框架',
            link: 'https://react.dev',
          },
          {
            label: 'Modern.js - 路由管理',
            link: 'https://modernjs.dev/',
          },
          {
            label: 'MUI - 组件库',
            link: 'https://mui.com/material-ui/',
          },
          {
            label: 'notistack - 通知管理',
            link: 'https://notistack.com/',
          },
          {
            label: 'Recoil - 状态管理',
            link: 'https://recoiljs.org/',
          },
          {
            label: 'Formik - 表单管理',
            link: 'https://formik.org/',
          },
          {
            label: 'SWR - 请求管理',
            link: 'https://swr.bootcss.com/',
          },
          {
            label: 'OpenAPI TypeScript - OpenAPI Client CodeGen 工具',
            link: 'https://openapi-ts.dev/',
          },
          {
            label: 'dayjs - 时间工具库',
            link: 'https://day.js.org/zh-CN/',
          },
        ]}
      />
      <DependenciesGroup
        label="Server"
        items={[
          {
            label: 'Echo - Go 语言服务端框架',
            link: 'https://echo.labstack.com/',
          },
          {
            label: 'GORM - Go 语言ORM',
            link: 'https://gorm.io/',
          },
          {
            label: 'oapi codegen - Go 语言 OpenAPI CodeGen 工具',
            link: 'https://github.com/oapi-codegen/oapi-codegen',
          },
          {
            label: 'Kubernetes client-go - Go 语言 Kubernetes 客户端',
            link: 'https://github.com/kubernetes/client-go',
          },
          {
            label: 'golang jwt - Go 语言 JSON Web Token 库',
            link: 'https://github.com/golang-jwt/jwt',
          },
        ]}
      />
      <DependenciesGroup
        label="其他主要依赖项"
        items={[
          { label: 'PostgreSQL', link: 'https://www.postgresql.org/' },
          {
            label: 'TypeSpec - 基于 TypeScript 语法的 IDL',
            link: 'https://typespec.io/',
          },
        ]}
      />
    </>
  );
}
