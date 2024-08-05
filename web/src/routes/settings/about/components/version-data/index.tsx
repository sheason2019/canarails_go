import { Typography, Button } from '@mui/material';
import useVersion from './hooks/use-version';
import useVersionUpdate from './hooks/use-version-update';

export default function VersionData() {
  const { data } = useVersion();
  const versionData = data.data;

  const { checkUpdate } = useVersionUpdate();

  return (
    <>
      <Typography>镜像摘要：{versionData?.gitHash}</Typography>
      <Typography>构建时间：{versionData?.buildTime}</Typography>

      <Button variant="contained" sx={{ mt: 1 }} onClick={() => checkUpdate()}>
        检查更新
      </Button>
    </>
  );
}
