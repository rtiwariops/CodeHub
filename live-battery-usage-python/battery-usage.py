import psutil
import time
import json

while True:
    battery = psutil.sensors_battery()
    battery_data = {}

    battery_data["battery_percentage"] = battery.percent
    battery_data["power_plugged_in"] = battery.power_plugged

    # battery.secsleft provides the approximate number of seconds left before the battery runs out
    if battery.secsleft == psutil.POWER_TIME_UNKNOWN:
        battery_data["battery_time_left"] = None
    elif battery.secsleft == psutil.POWER_TIME_UNLIMITED:
        battery_data["battery_time_left"] = "Battery is currently being charged."
    else:
        # converting seconds to hours, minutes, and seconds
        hours_left = battery.secsleft // 3600
        minutes_left = (battery.secsleft % 3600) // 60
        seconds_left = (battery.secsleft % 60)
        battery_data["battery_time_left"] = {"hours": hours_left, "minutes": minutes_left, "seconds": seconds_left}

    # Convert dictionary to JSON and print
    print(json.dumps(battery_data, indent=4))

    # Sleep for a minute before fetching the information again
    time.sleep(60)