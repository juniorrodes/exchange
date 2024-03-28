package model

var (
    CurrenciesCodes = []string{"USD", "EUR", "JPY", "GBP", "AUD", "CAD", "CHF", "CNY", "SEK",
        "NZD", "BRL", "RUB", "INR", "TRY", "KRW", "ZAR", "HKD", "SGD", "NOK", "MXN", "IDR",
        "MYR", "PLN", "DKK", "THB", "PHP", "HUF", "CZK", "ILS", "CLP", "PKR", "AED", "SAR",
        "TWD", "KWD", "QAR", "EGP", "NPR", "BDT", "OMR", "COP", "ARS", "VND", "IQD", "JOD",
        "KES", "UGX", "BYN", "LKR", "CUP", "GHS", "TZS", "SYP", "MMK", "UZS", "SDG", "KHR",
        "MAD", "GNF", "RSD", "YER", "MZN", "AFN", "NAD", "AZN", "ALL", "DZD", "AOA", "XOF",
        "XAF", "XCD", "XPF", "ANG", "AWG", "BBD", "BSD", "BMD", "BHD", "BZD", "BWP", "BND",
        "BIF", "CVE", "KYD", "XAF", "CUC", "DJF", "FJD", "GMD", "GYD", "HTG", "HNL", "JMD",
        "KZT", "KWD", "KGS", "LAK", "LSL", "LRD", "LYD", "MOP", "MWK", "MVR", "MRO", "MUR",
        "MDL", "MNT", "MAD", "MOP", "MRO", "MUR", "MZN", "NAD", "NPR", "NIO", "NGN", "KPW",
        "OMR", "PGK", "PYG", "RWF", "WST", "STD", "SLL", "SOS", "SRD", "SZL", "TJS", "TOP",
        "TTD", "TND", "TMT", "UGX", "UAH", "AED", "VUV", "VUV", "WST", "YER", "ZMW", "ZWL",
    }
)

type Currencies map[string]map[string]float32
