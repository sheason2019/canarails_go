import { Box, Stack, Typography } from '@mui/material';
import useAppVariant from './hooks/use-app-variant';
import Descriptions from '@/common/descriptions';
import AddHeaderMatchButton from './components/add-header-match-button';
import PutInfoButton from './components/put-info-button';
import PutConfigButton from './components/put-config-button';
import HeaderMatches from './components/header-matches';
import HeaderMatchAlert from './components/header-matches/header-match-alert';

export default function Page() {
  const { data } = useAppVariant();

  const appVar = data?.data;

  return (
    <main>
      <Stack sx={{ my: 2 }} direction="row">
        <Typography variant="h5" sx={{ flex: 1 }}>
          基本信息
        </Typography>
        <PutInfoButton />
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
        <PutConfigButton />
      </Stack>
      <Descriptions
        items={[
          {
            label: '映射端口',
            value: appVar?.exposePort,
          },
          {
            label: '镜像名称',
            value: appVar?.imageName,
          },
          {
            label: '实例数量',
            value: appVar?.replicas,
          },
        ]}
      />
      <Stack sx={{ my: 2 }} direction="row">
        <Typography variant="h5" sx={{ flex: 1 }}>
          标头匹配
        </Typography>
        <AddHeaderMatchButton />
      </Stack>
      <HeaderMatchAlert />
      <HeaderMatches />
    </main>
  );
}
