import { Outlet, useLocation, useNavigate } from '@modern-js/runtime/router';
import { Box, Tab, Tabs, Typography } from '@mui/material';
import useAppVariant from './hooks/use-app-variant';
import { useMemo } from 'react';

export default function Layout() {
  const { data } = useAppVariant();
  const appVar = data?.data;
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
        <Typography variant="h4">{appVar?.title}</Typography>
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
              navigate(
                `/apps/${appVar?.appId}/app-variants/${appVar?.id}/logs`,
              )
            }
          />
        </Tabs>
      </Box>
      <Outlet />
    </div>
  );
}
