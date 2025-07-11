<script setup lang="ts">
import { reactive, ref, watch } from 'vue';

// Define Prop Types
interface FormData {
  [key: string]: any; // Memungkinkan properti dinamis
  name: string;
  quantity: number;
}

interface Errors {
  [key: string]: string | undefined;
}

// Define Props
const props = withDefaults(defineProps<{
  title?: string;
  initialData: FormData;
  submitButtonText?: string;
  validationRules: {
    [key: string]: (value: any) => string | undefined; // Fungsi validasi mengembalikan pesan error atau undefined
  };
  onSubmit: (data: FormData) => Promise<void>;
}>(), {
  title: 'Form Example',
  submitButtonText: 'Submit',
});

// Reactive State
const formData = reactive<FormData>({ ...props.initialData });
const errors = reactive<Errors>({});
const isSubmitting = ref<boolean>(false);
const submissionError = ref<string | null>(null);

// Watch untuk memperbarui formData jika initialData berubah (misal saat edit)
watch(() => props.initialData, (newData: FormData) => {
  Object.assign(formData, newData);
}, { deep: true });

// Field Validation
const validateField = (field: string) => {
  if (props.validationRules[field]) {
    const rule = props.validationRules[field];
    const errorMessage = rule(formData[field]);
    if (errorMessage) {
      errors[field] = errorMessage;
    } else {
      delete errors[field];
    }
  }
};

// Form Validation
const validateForm = (): boolean => {
  let isValid = true;
  for (const field in props.validationRules) {
    validateField(field);
    if (errors[field]) {
      isValid = false;
    }
  }
  return isValid;
};

// Handle Form Submission
const handleSubmit = async () => {
  submissionError.value = null;
  if (!validateForm()) {
    return;
  }

  isSubmitting.value = true;
  try {
    await props.onSubmit(formData);
    // Jika berhasil, form bisa direset atau ada navigasi
    // Object.assign(formData, props.initialData); // Reset form jika diperlukan
  } catch (err: any) {
    submissionError.value = err.message || 'An unexpected error occurred during submission.';
  } finally {
    isSubmitting.value = false;
  }
};
</script>


<template>
    <form @submit.prevent="handleSubmit" class="space-y-4 p-6 bg-white shadow rounded-lg">
      <h2 class="text-2xl font-bold text-gray-800 mb-6">{{ title }}</h2>

      <div>
        <label for="name" class="block text-sm font-medium text-gray-700">Name</label>
        <input
          type="text"
          id="name"
          v-model="formData.name"
          @blur="validateField('name')"
          class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm p-2 focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          :class="{ 'border-red-500': errors.name }"
        />
        <p v-if="errors.name" class="mt-1 text-sm text-red-600">{{ errors.name }}</p>
      </div>

      <div>
        <label for="quantity" class="block text-sm font-medium text-gray-700">Quantity</label>
        <input
          type="number"
          id="quantity"
          v-model.number="formData.quantity"
          @blur="validateField('quantity')"
          class="mt-1 block w-full border border-gray-300 rounded-md shadow-sm p-2 focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          :class="{ 'border-red-500': errors.quantity }"
        />
        <p v-if="errors.quantity" class="mt-1 text-sm text-red-600">{{ errors.quantity }}</p>
      </div>

      <slot name="additional-fields" :formData="formData" :errors="errors"></slot>

      <div v-if="submissionError" class="text-red-600 text-sm mb-4">{{ submissionError }}</div>

      <button
        type="submit"
        class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        :disabled="isSubmitting"
      >
        {{ isSubmitting ? 'Submitting...' : submitButtonText }}
      </button>
    </form>
  </template>

