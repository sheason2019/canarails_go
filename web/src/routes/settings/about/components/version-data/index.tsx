import { Typography } from '@mui/material';
import useVersion from './hooks/use-version';

export default function VersionData() {
  const { data } = useVersion();
  const versionData = data.data;

  return (
    <>
      <Typography>镜像摘要：{versionData?.gitHash}</Typography>
      <Typography>构建时间：{versionData?.buildTime}</Typography>
    </>
  );
}
