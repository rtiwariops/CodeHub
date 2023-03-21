import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import accuracy_score, confusion_matrix
import matplotlib.pyplot as plt
import seaborn as sns

# Load and preprocess data (age, income, distance, missed_appointments)
data = pd.read_csv("patient_data.csv")
X = data.drop("missed_appointments", axis=1)
y = data["missed_appointments"]

X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)

# Train a RandomForestClassifier
clf = RandomForestClassifier(random_state=42)
clf.fit(X_train, y_train)

# Predict on test data
y_pred = clf.predict(X_test)

# Calculate accuracy
accuracy = accuracy_score(y_test, y_pred)
print(f"Accuracy: {accuracy}")

# Print the actual and predicted values
print("Actual values:", list(y_test))
print("Predicted values:", list(y_pred))

# Print a summary of the results
print(f"\nOut of {len(y_test)} test cases, the model correctly predicted {int(accuracy * len(y_test))} cases, with an accuracy of {accuracy * 100}%.")

# Calculate feature importances
importances = clf.feature_importances_
feature_names = X.columns

# Plot feature importances
plt.bar(feature_names, importances)
plt.xlabel("Features")
plt.ylabel("Importance")
plt.title("Feature Importance in Predicting Missed Appointments")
plt.show()

# Calculate confusion matrix
cm = confusion_matrix(y_test, y_pred)

# Plot confusion matrix heatmap
plt.figure(figsize=(6, 4))
sns.heatmap(cm, annot=True, cmap="Blues", fmt="d", cbar=False, xticklabels=['Not Missed', 'Missed'], yticklabels=['Not Missed', 'Missed'])
plt.xlabel("Predicted")
plt.ylabel("Actual")
plt.title("Confusion Matrix for Missed Appointment Predictions")
plt.show()
