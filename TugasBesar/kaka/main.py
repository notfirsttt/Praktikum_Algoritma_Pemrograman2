import numpy as np
from PIL import Image
import requests
import time
from io import BytesIO

def compress_recursive(image, threshold):
    return compress_helper(image, 0, 0, image.shape[1], image.shape[0], threshold)

def compress_helper(image, start_x, start_y, width, height, threshold):
    if width <= 1 or height <= 1:
        return image

    std_dev = calculate_standard_deviation(image, start_x, start_y, width, height)
    if std_dev <= threshold:
        average_color = calculate_average_color(image, start_x, start_y, width, height)
        fill_region(image, start_x, start_y, width, height, average_color)
        return image

    mid_x = start_x + width // 2
    mid_y = start_y + height // 2

    compress_helper(image, start_x, start_y, width // 2, height // 2, threshold)  # Top-left
    compress_helper(image, mid_x, start_y, width // 2, height // 2, threshold)  # Top-right
    compress_helper(image, start_x, mid_y, width // 2, height // 2, threshold)  # Bottom-left
    compress_helper(image, mid_x, mid_y, width // 2, height // 2, threshold)  # Bottom-right

    return image

def compress_iterative(image, threshold):
    step = 2
    while step <= min(image.shape[1], image.shape[0]):
        for x in range(0, image.shape[1], step):
            for y in range(0, image.shape[0], step):
                width = min(step, image.shape[1] - x)
                height = min(step, image.shape[0] - y)

                std_dev = calculate_standard_deviation(image, x, y, width, height)
                if std_dev <= threshold:
                    average_color = calculate_average_color(image, x, y, width, height)
                    fill_region(image, x, y, width, height, average_color)
        step *= 2
    return image

def calculate_standard_deviation(image, start_x, start_y, width, height):
    region = image[start_y:start_y + height, start_x:start_x + width]
    return np.std(region)

def calculate_average_color(image, start_x, start_y, width, height):
    region = image[start_y:start_y + height, start_x:start_x + width]
    return int(np.mean(region))

def fill_region(image, start_x, start_y, width, height, color):
    image[start_y:start_y + height, start_x:start_x + width] = color

def download_image(url):
    response = requests.get(url)
    return Image.open(BytesIO(response.content))

def get_image_size(image):
    buffer = BytesIO()
    image.save(buffer, format="JPEG")
    return len(buffer.getvalue()) // 1024  # Size in KB

def main():
    image_urls = [
        "https://lh5.googleusercontent.com/p/AF1QipOtn45cWWNc4uyT2hJiwTFZ7U5Qo6FI2o_jzWfS=w810-h468-n-k-no",
        "https://lh5.googleusercontent.com/p/AF1QipMI2ZIJLjw1pHZTB5bvdwNvcnUWn-WbbMblfJFz=w810-h468-n-k-no",
        "https://encrypted-tbn2.gstatic.com/licensed-image?q=tbn:ANd9GcSynVvyP9oLtdZoH2_mD6AWA6oPRGJTcM82ePsQ3Y6dGntZy0RaBx4N3bMVUiDwRrfewII4YZ5Ii9eIl5Y4943_pLlrdoBA8en0qtaPBw",
        "https://lh5.googleusercontent.com/p/AF1QipNCFJOlcUwP6eoSmvOWPVOqY6eI1_TWuRTkGRME=w810-h468-n-k-no"
    ]

    threshold = 10

    print("+-------------------+-------------------------+------------------------+-------------------------+------------------------+----------------------+")
    print("| Image             | Original Size (KB)      | Iterative Compression  | Iterative Time (s)      | Recursive Compression | Recursive Time (s)   |")
    print("+-------------------+-------------------------+------------------------+-------------------------+------------------------+----------------------+")

    for idx, image_url in enumerate(image_urls, start=1):
        try:
            original_image = download_image(image_url)
            image_array = np.array(original_image.convert("L"))  # Convert to grayscale

            original_size_kb = get_image_size(original_image)

            # Iterative compression
            start_time = time.time()
            iterative_result = compress_iterative(image_array.copy(), threshold)
            iterative_time = time.time() - start_time

            iterative_image = Image.fromarray(iterative_result)
            iterative_size_kb = get_image_size(iterative_image)

            # Recursive compression
            start_time = time.time()
            recursive_result = compress_recursive(image_array.copy(), threshold)
            recursive_time = time.time() - start_time

            recursive_image = Image.fromarray(recursive_result)
            recursive_size_kb = get_image_size(recursive_image)

            print(f"| Gambar {idx:<14} | {original_size_kb:<23} | {iterative_size_kb:<22} | {iterative_time:<22.10f} | {recursive_size_kb:<22} | {recursive_time:<22.10f} |")
            print("+-------------------+-------------------------+------------------------+-------------------------+------------------------+----------------------+")

        except Exception as e:
            print(f"Error processing image {idx}: {e}")

if __name__ == "__main__":
    main()
