import { Outlet, useLocation, useNavigate } from '@modern-js/runtime/router';
import { Box, Stack, Tab, Tabs, Typography } from '@mui/material';
import useAppVariant from './hooks/use-app-variant';
import { useMemo } from 'react';
import useApp from '../apps/[id]/hooks/use-app';
import StyledLink from '@/common/styled-link';

export default function Layout() {
  const { data: appData } = useApp();
  const app = appData?.data;

  const { data: appVarData } = useAppVariant();
  const appVar = appVarData?.data;

  const navigate = useNavigate();
  const location = useLocation();

  const tabValue = useMemo(() => {
    const tabName = location.pathname.split('/')[5];
    switch (tabName) {
      case 'logs':
        return 1;
      default:
        return 0;
    }
  }, [location.pathname]);

  return (
    <div>
      <Box sx={{ my: 2 }}>
        <Stack direction="row" alignItems="baseline">
          <Typography
            component={StyledLink}
            to={`/apps/${app?.id}`}
            sx={{ textDecoration: 'none' }}
            color="inherit"
            variant="h4"
          >
            {app?.title}
          </Typography>
          <Typography sx={{ fontSize: 28, mx: 1 }}>{`/`}</Typography>
          <Typography variant="h5" color="GrayText">
            {appVar?.title}
          </Typography>
        </Stack>
        <Typography variant="body2" color="GrayText">
          ID {appVar?.id}
        </Typography>
      </Box>
      <Box sx={{ borderBottom: 1, borderColor: 'divider' }}>
        <Tabs value={tabValue}>
          <Tab
            label="泳道信息"
            onClick={() =>
              navigate(`/apps/${appVar?.appId}/app-variants/${appVar?.id}`)
            }
          />
          <Tab
            label="日志信息"
            onClick={() =>
              navigate(`/apps/${appVar?.appId}/app-variants/${appVar?.id}/logs`)
            }
          />
        </Tabs>
      </Box>
      <Outlet />
    </div>
  );
}
