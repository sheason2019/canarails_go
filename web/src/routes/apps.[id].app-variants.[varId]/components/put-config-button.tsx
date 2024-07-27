import useApi from '@/common/use-api';
import useAppVariant from '../hooks/use-app-variant';
import { useSnackbar } from 'notistack';
import useDialog from '@/common/use-dialog';
import { useFormik } from 'formik';
import * as Yup from 'yup';
import useFieldProperties from '@/routes/hooks/use-field-properties';
import { useEffect } from 'react';
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Stack,
  TextField,
} from '@mui/material';

export default function PutConfigButton() {
  const { data, mutate } = useAppVariant();
  const appVar = data?.data;

  const api = useApi();
  const { enqueueSnackbar } = useSnackbar();
  const { isOpen, onClose, onOpen } = useDialog();

  const formik = useFormik({
    initialValues: {
      exposePort: appVar?.exposePort ?? 80,
      imageName: appVar?.imageName ?? '',
      replicas: appVar?.replicas ?? 0,
    },
    validationSchema: Yup.object({
      exposePort: Yup.number()
        .required('映射端口不能为空')
        .integer('映射端口必须为正整数')
        .min(0, '映射端口必须为正整数'),
      imageName: Yup.string(),
      replicas: Yup.number()
        .required('实例数量不能为空')
        .integer('实例数量必须为正整数')
        .min(0, '实例数量必须为正整数'),
    }),
    onSubmit: async values => {
      if (!appVar) throw new Error('cannot find appVar');

      await api.PUT('/api/app-variant/{id}', {
        params: {
          path: { id: appVar.id },
        },
        body: {
          ...appVar,
          ...values,
        },
      });
      enqueueSnackbar('修改配置信息成功', { variant: 'success' });
      onClose();
      mutate();
    },
  });

  const { getFieldProperties } = useFieldProperties(formik);

  useEffect(() => {
    formik.resetForm();
  }, [isOpen]);

  return (
    <>
      <Button variant="contained" onClick={onOpen}>
        修改
      </Button>
      <Dialog open={isOpen} onClose={onClose} fullWidth>
        <form onSubmit={formik.handleSubmit}>
          <DialogTitle>修改配置信息</DialogTitle>
          <DialogContent>
            <Stack spacing={2}>
              <TextField
                label="映射端口"
                variant="standard"
                type="number"
                {...getFieldProperties('exposePort')}
              />
              <TextField
                label="镜像名称"
                variant="standard"
                {...getFieldProperties('imageName')}
              />
              <TextField
                label="实例数量"
                variant="standard"
                type="number"
                {...getFieldProperties('replicas')}
              />
            </Stack>
          </DialogContent>
          <DialogActions>
            <Button onClick={onClose}>取消</Button>
            <Button variant="contained" type="submit">
              提交
            </Button>
          </DialogActions>
        </form>
      </Dialog>
    </>
  );
}
