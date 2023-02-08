<!-- @format -->

# System Monitoring 🖥️

A simple Python script to monitor the system's kernel, memory, and CPU usage.

## Requirements 📋

Python 3.x
psutil module

## Usage 💻

Run the script with the following command:

```
python system_monitoring.py
```

## Output 📈

```
========================================
System Monitoring
========================================
{
    "kernel_info": {
        "kernel_version": "5.11.0-rc1-mainline",
        "system_name": "Linux",
        "node_name": "ubuntu",
        "machine": "x86_64"
    },
    "memory_info": {
        "total_memory": 7.450580596923828,
        "available_memory": 4.557059288024902,
        "used_memory": 2.893521385192871,
        "memory_percentage": 62.0
    },
    "cpu_info": {
        "physical_cores": 2,
        "total_cores": 4,
        "processor_speed": 2901.0,
        "cpu_usage_per_core": {
            0: 1.0,
            1: 0.0,
            2: 0.0,
            3: 0.0
        },
        "total_cpu_usage": 1.0
    }
}
```

## Explanation 💡

- kernel_info: Information about the kernel version, system name, node name, and machine type.
- memory_info: Information about the total memory, available memory, used memory, and memory usage percentage.
- cpu_info: Information about the physical cores, total cores, processor speed, CPU usage per core, and total CPU usage.

## Note 📝

The output may vary based on the system's configuration.
