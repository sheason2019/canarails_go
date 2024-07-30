import useApp from '@/routes/apps/[id]/hooks/use-app';
import useAppVariant from '../../hooks/use-app-variant';
import { Alert } from '@mui/material';

export default function HeaderMatchAlert() {
  const { data: appData } = useApp();
  const app = appData?.data;

  const { data: appVarData } = useAppVariant();
  const appVar = appVarData?.data;

  if (appVar?.id === app?.defaultVariantId) {
    return (
      <Alert>
        当前流量泳道为默认泳道，当用户流量未匹配到其他泳道时，流量将被转发到该泳道
      </Alert>
    );
  }

  if (!appVar?.matches.length) {
    return (
      <Alert severity="warning">
        当前流量泳道不是默认泳道，且没有配置标头匹配，这意味着流量泳道配置不会应用到
        Kubernetes 集群
      </Alert>
    );
  }

  return null;
}
