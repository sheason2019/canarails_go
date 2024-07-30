import { Stack, Typography } from '@mui/material';
import useApp from './hooks/use-app';
import Descriptions from '@/common/descriptions';
import AddAppHostnameButton from './components/add-app-hostname-button';
import AppHostnames from './components/app-hostnames';
import PutAppInfoButton from './components/put-app-info-button';

export default function AppPage() {
  const { data } = useApp();
  const app = data?.data;

  return (
    <main>
      <Stack
        justifyContent="space-between"
        alignItems="center"
        direction="row"
        sx={{ my: 2 }}
      >
        <Typography variant="h5">基本信息</Typography>
        <PutAppInfoButton />
      </Stack>
      <Descriptions
        items={[
          {
            label: 'ID',
            value: app?.id,
          },
          {
            label: 'App 名称',
            value: app?.title,
          },
          { label: 'App 简介', value: app?.description },
        ]}
      />
      <Stack
        justifyContent="space-between"
        alignItems="center"
        direction="row"
        sx={{ my: 2 }}
      >
        <Typography variant="h5">域名匹配</Typography>
        <AddAppHostnameButton />
      </Stack>
      <AppHostnames />
    </main>
  );
}
