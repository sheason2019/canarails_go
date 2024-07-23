import { TextFieldProps } from '@mui/material';
import { useFormik } from 'formik';
import { ReactNode } from 'react';

export default function useFieldProperties<
  T extends ReturnType<typeof useFormik<any>>,
>(formik: T) {
  const getFieldProperties = (
    name: keyof T['initialValues'] & string,
  ): Partial<TextFieldProps> => ({
    name,
    value: formik.values[name],
    onChange: formik.handleChange,
    onBlur: formik.handleBlur,
    error: formik.touched[name] && Boolean(formik.errors[name]),
    helperText: formik.touched[name] && (formik.errors[name] as ReactNode),
    disabled: formik.isSubmitting,
  });

  return {
    getFieldProperties,
  };
}
