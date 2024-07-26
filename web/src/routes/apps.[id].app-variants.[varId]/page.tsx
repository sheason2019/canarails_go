import { Box, Button, Stack, Typography } from '@mui/material';
import useAppVariant from './hooks/use-app-variant';
import Descriptions from '@/common/descriptions';

export default function Page() {
  const { data } = useAppVariant();

  const appVar = data?.data;

  return (
    <main>
      <Stack sx={{ my: 2 }} direction="row">
        <Typography variant="h5" sx={{ flex: 1 }}>
          基本信息
        </Typography>
        <Button variant="contained">修改</Button>
      </Stack>
      <Descriptions
        items={[
          {
            label: 'ID',
            value: appVar?.id,
          },
          {
            label: '泳道名称',
            value: appVar?.title,
          },
          {
            label: '简介',
            value: appVar?.description,
          },
        ]}
      />
      <Stack sx={{ my: 2 }} direction="row">
        <Typography variant="h5" sx={{ flex: 1 }}>
          配置信息
        </Typography>
        <Button variant="contained">修改</Button>
      </Stack>
      <Descriptions
        items={[
          {
            label: '映射端口',
            value: appVar?.exposePort,
          },
          {
            label: '镜像名称',
            value: '',
          },
          {
            label: '实例数量',
            value: '',
          },
        ]}
      />
      <Stack sx={{ my: 2 }} direction="row">
        <Typography variant="h5" sx={{ flex: 1 }}>
          标头匹配
        </Typography>
        <Button variant="contained">新增</Button>
      </Stack>
      {!!appVar?.matches.length ? (
        <Descriptions
          items={
            appVar?.matches.map(match => ({
              label: match.header,
              value: match.value,
            })) ?? []
          }
        />
      ) : (
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            height: 48,
          }}
        >
          <Typography variant="body2" color="GrayText">
            暂无数据
          </Typography>
        </Box>
      )}
    </main>
  );
}
