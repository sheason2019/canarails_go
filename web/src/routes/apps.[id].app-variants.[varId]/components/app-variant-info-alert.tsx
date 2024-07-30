import { Alert } from '@mui/material';
import useAppVariant from '../hooks/use-app-variant';

export default function AppVariantInfoAlert() {
  const { data } = useAppVariant();
  const appVar = data.data;

  if (!appVar?.imageName.length || appVar.replicas <= 0) {
    return (
      <Alert severity="warning">
        镜像名称或实例数量不正确，流量泳道配置将不会正确生效
      </Alert>
    );
  }

  return null;
}
