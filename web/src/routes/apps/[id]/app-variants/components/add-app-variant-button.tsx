import useApi from '@/common/use-api';
import useDialog from '@/common/use-dialog';
import useFieldProperties from '@/routes/hooks/use-field-properties';
import { useParams } from '@modern-js/runtime/router';
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Stack,
  TextField,
} from '@mui/material';
import { useFormik } from 'formik';
import { useSnackbar } from 'notistack';
import { useEffect } from 'react';
import * as Yup from 'yup';
import useAppVariants from './app-variants-table/hooks/use-app-variants';

export default function AddAppVariantButton() {
  const { id } = useParams();
  const api = useApi();
  const { isOpen, onOpen, onClose } = useDialog();
  const { enqueueSnackbar } = useSnackbar();
  const { mutate } = useAppVariants();

  const formik = useFormik({
    initialValues: {
      title: '',
      description: '',
    },
    validationSchema: Yup.object({
      title: Yup.string().required('泳道名称不能为空'),
      Description: Yup.string(),
    }),
    onSubmit: async values => {
      const res = await api.POST('/api/app-variant', {
        body: {
          title: values.title,
          description: values.description,
          appId: Number(id),
          exposePort: 80,
        },
      });
      enqueueSnackbar(`创建流量泳道成功，ID：${res.data}`, {
        variant: 'success',
      });
      mutate();
      onClose();
    },
  });

  const { getFieldProperties } = useFieldProperties(formik);

  useEffect(() => {
    formik.resetForm();
  }, [isOpen]);

  return (
    <>
      <Button variant="contained" onClick={onOpen}>
        新增
      </Button>
      <Dialog open={isOpen} onClose={onClose} fullWidth>
        <form onSubmit={formik.handleSubmit}>
          <DialogTitle>创建流量泳道</DialogTitle>
          <DialogContent>
            <Stack spacing={2}>
              <TextField
                label="泳道名称"
                variant="standard"
                {...getFieldProperties('title')}
              />
              <TextField
                label="简介"
                variant="standard"
                multiline
                minRows={3}
                maxRows={5}
                {...getFieldProperties('description')}
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
