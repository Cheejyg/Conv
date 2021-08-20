package com.github.cheejyg.conv.java;

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import com.github.cheejyg.conv.java.*;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Conversion.
 * 
 * @author Cheejyg
 */
public class Conv {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// ToString
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ToString parses <code>Object</code> object_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param object_ The <code>Object</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static java.lang.String ToString(java.lang.Object object_) {
		return object_.toString();
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// String conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * StringToBoolean parses <code>String</code> string_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param string_ The <code>String</code> to be parsed.
	 * @return <code>Boolean</code>
	 * @see com.github.cheejyg.conv.java.String
	 * @see com.github.cheejyg.conv.java.String#StringToBoolean(java.lang.String
	 *      string_)
	 */
	public static java.lang.Boolean StringToBoolean(java.lang.String string_) {
		return String.StringToBoolean(string_);
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Array conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ObjectsToObjectArray parses <code>Object[]</code> objects_ to
	 * <code>Object[]</code>.
	 * 
	 * @author Cheejyg
	 * @param objects_ The <code>Object[]</code> to be parsed.
	 * @return <code>Object[]</code>
	 * @see com.github.cheejyg.conv.java.Array
	 * @see com.github.cheejyg.conv.java.Array#ObjectsToObjectArray(java.lang.Object...
	 *      objects_)
	 */
	public static java.lang.Object[] ObjectsToObjectArray(java.lang.Object... objects_) {
		return Array.ObjectsToObjectArray(objects_);
	}

	/**
	 * LongsToLongArray parses <code>Long[]</code> longs_ to <code>Long[]</code>.
	 * 
	 * @author Cheejyg
	 * @param longs_ The <code>Long[]</code> to be parsed.
	 * @return <code>Long[]</code>
	 * @see com.github.cheejyg.conv.java.Array
	 * @see com.github.cheejyg.conv.java.Array#LongsToLongArray(java.lang.Long...
	 *      longs_)
	 */
	public static java.lang.Long[] LongsToLongArray(java.lang.Long... longs_) {
		return Array.LongsToLongArray(longs_);
	}

	/**
	 * IntegersToIntegerArray parses <code>Integer[]</code> integers_ to
	 * <code>Integer[]</code>.
	 * 
	 * @author Cheejyg
	 * @param integers_ The <code>Integer[]</code> to be parsed.
	 * @return <code>Integer[]</code>
	 * @see com.github.cheejyg.conv.java.Array
	 * @see com.github.cheejyg.conv.java.Array#IntegersToIntegerArray(java.lang.Integer...
	 *      integers_)
	 */
	public static java.lang.Integer[] IntegersToIntegerArray(java.lang.Integer... integers_) {
		return Array.IntegersToIntegerArray(integers_);
	}

	/**
	 * ShortsToShortArray parses <code>Short[]</code> shorts_ to
	 * <code>Short[]</code>.
	 * 
	 * @author Cheejyg
	 * @param shorts_ The <code>Short[]</code> to be parsed.
	 * @return <code>Short[]</code>
	 * @see com.github.cheejyg.conv.java.Array
	 * @see com.github.cheejyg.conv.java.Array#ShortsToShortArray(java.lang.Short...
	 *      shorts_)
	 */
	public static java.lang.Short[] ShortsToShortArray(java.lang.Short... shorts_) {
		return Array.ShortsToShortArray(shorts_);
	}

	/**
	 * BytesToByteArray parses <code>Byte[]</code> bytes_ to <code>Byte[]</code>.
	 * 
	 * @author Cheejyg
	 * @param bytes_ The <code>Byte[]</code> to be parsed.
	 * @return <code>Byte[]</code>
	 * @see com.github.cheejyg.conv.java.Array
	 * @see com.github.cheejyg.conv.java.Array#BytesToByteArray(java.lang.Byte...
	 *      bytes_)
	 */
	public static java.lang.Byte[] BytesToByteArray(java.lang.Byte... bytes_) {
		return Array.BytesToByteArray(bytes_);
	}

	/**
	 * BooleansToBooleanArray parses <code>Boolean[]</code> booleans_ to
	 * <code>Boolean[]</code>.
	 * 
	 * @author Cheejyg
	 * @param booleans_ The <code>Boolean[]</code> to be parsed.
	 * @return <code>Boolean[]</code>
	 * @see com.github.cheejyg.conv.java.Array
	 * @see com.github.cheejyg.conv.java.Array#BooleansToBooleanArray(java.lang.Boolean...
	 *      booleans_)
	 */
	public static java.lang.Boolean[] BooleansToBooleanArray(java.lang.Boolean... booleans_) {
		return Array.BooleansToBooleanArray(booleans_);
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Long conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * LongToString parses <code>long</code> long_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>String</code>
	 * @see com.github.cheejyg.conv.java.Long
	 * @see com.github.cheejyg.conv.java.Long#LongToString(long long_)
	 */
	public static java.lang.String LongToString(long long_) {
		return Long.LongToString(long_);
	}

	/**
	 * LongToLong parses <code>long</code> long_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Long</code>
	 * @see com.github.cheejyg.conv.java.Long
	 * @see com.github.cheejyg.conv.java.Long#LongToLong(long long_)
	 */
	public static java.lang.Long LongToLong(long long_) {
		return Long.LongToLong(long_);
	}

	/**
	 * LongToInteger parses <code>long</code> long_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Integer</code>
	 * @see com.github.cheejyg.conv.java.Long
	 * @see com.github.cheejyg.conv.java.Long#LongToInteger(long long_)
	 */
	public static java.lang.Integer LongToInteger(long long_) {
		return Long.LongToInteger(long_);
	}

	/**
	 * LongToShort parses <code>long</code> long_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Short</code>
	 * @see com.github.cheejyg.conv.java.Long
	 * @see com.github.cheejyg.conv.java.Long#LongToShort(long long_)
	 */
	public static java.lang.Short LongToShort(long long_) {
		return Long.LongToShort(long_);
	}

	/**
	 * LongToByte parses <code>long</code> long_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Byte</code>
	 * @see com.github.cheejyg.conv.java.Long
	 * @see com.github.cheejyg.conv.java.Long#LongToByte(long long_)
	 */
	public static java.lang.Byte LongToByte(long long_) {
		return Long.LongToByte(long_);
	}

	/**
	 * LongToBoolean parses <code>long</code> long_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Boolean</code>
	 * @see com.github.cheejyg.conv.java.Long
	 * @see com.github.cheejyg.conv.java.Long#LongToBoolean(long long_)
	 */
	public static java.lang.Boolean LongToBoolean(long long_) {
		return Long.LongToBoolean(long_);
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Int conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * IntToString parses <code>int</code> int_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>String</code>
	 * @see com.github.cheejyg.conv.java.Int
	 * @see com.github.cheejyg.conv.java.Int#IntToString(int int_)
	 */
	public static java.lang.String IntToString(int int_) {
		return Int.IntToString(int_);
	}

	/**
	 * IntToLong parses <code>int</code> int_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Long</code>
	 * @see com.github.cheejyg.conv.java.Int
	 * @see com.github.cheejyg.conv.java.Int#IntToLong(int int_)
	 */
	public static java.lang.Long IntToLong(int int_) {
		return Int.IntToLong(int_);
	}

	/**
	 * IntToInteger parses <code>int</code> int_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Integer</code>
	 * @see com.github.cheejyg.conv.java.Int
	 * @see com.github.cheejyg.conv.java.Int#IntToInteger(int int_)
	 */
	public static java.lang.Integer IntToInteger(int int_) {
		return Int.IntToInteger(int_);
	}

	/**
	 * IntToShort parses <code>int</code> int_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Short</code>
	 * @see com.github.cheejyg.conv.java.Int
	 * @see com.github.cheejyg.conv.java.Int#IntToShort(int int_)
	 */
	public static java.lang.Short IntToShort(int int_) {
		return Int.IntToShort(int_);
	}

	/**
	 * IntToByte parses <code>int</code> int_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Byte</code>
	 * @see com.github.cheejyg.conv.java.Int
	 * @see com.github.cheejyg.conv.java.Int#IntToByte(int int_)
	 */
	public static java.lang.Byte IntToByte(int int_) {
		return Int.IntToByte(int_);
	}

	/**
	 * IntToBoolean parses <code>int</code> int_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param int_ The <code>int</code> to be parsed.
	 * @return <code>Boolean</code>
	 * @see com.github.cheejyg.conv.java.Int
	 * @see com.github.cheejyg.conv.java.Int#IntToBoolean(int int_)
	 */
	public static java.lang.Boolean IntToBoolean(int int_) {
		return Int.IntToBoolean(int_);
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Short conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ShortToString parses <code>short</code> short_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>String</code>
	 * @see com.github.cheejyg.conv.java.Short
	 * @see com.github.cheejyg.conv.java.Short#ShortToString(short short_)
	 */
	public static java.lang.String ShortToString(short short_) {
		return Short.ShortToString(short_);
	}

	/**
	 * ShortToLong parses <code>short</code> short_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Long</code>
	 * @see com.github.cheejyg.conv.java.Short
	 * @see com.github.cheejyg.conv.java.Short#ShortToLong(short short_)
	 */
	public static java.lang.Long ShortToLong(short short_) {
		return Short.ShortToLong(short_);
	}

	/**
	 * ShortToInteger parses <code>short</code> short_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Integer</code>
	 * @see com.github.cheejyg.conv.java.Short
	 * @see com.github.cheejyg.conv.java.Short#ShortToInteger(short short_)
	 */
	public static java.lang.Integer ShortToInteger(short short_) {
		return Short.ShortToInteger(short_);
	}

	/**
	 * ShortToShort parses <code>short</code> short_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Short</code>
	 * @see com.github.cheejyg.conv.java.Short
	 * @see com.github.cheejyg.conv.java.Short#ShortToShort(short short_)
	 */
	public static java.lang.Short ShortToShort(short short_) {
		return Short.ShortToShort(short_);
	}

	/**
	 * ShortToByte parses <code>short</code> short_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Byte</code>
	 * @see com.github.cheejyg.conv.java.Short
	 * @see com.github.cheejyg.conv.java.Short#ShortToByte(short short_)
	 */
	public static java.lang.Byte ShortToByte(short short_) {
		return Short.ShortToByte(short_);
	}

	/**
	 * ShortToBoolean parses <code>short</code> short_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param short_ The <code>short</code> to be parsed.
	 * @return <code>Boolean</code>
	 * @see com.github.cheejyg.conv.java.Short
	 * @see com.github.cheejyg.conv.java.Short#ShortToBoolean(short short_)
	 */
	public static java.lang.Boolean ShortToBoolean(short short_) {
		return Short.ShortToBoolean(short_);
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Byte conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * ByteToString parses <code>byte</code> byte_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>String</code>
	 * @see com.github.cheejyg.conv.java.Byte
	 * @see com.github.cheejyg.conv.java.Byte#ByteToString(byte byte_)
	 */
	public static java.lang.String ByteToString(byte byte_) {
		return Byte.ByteToString(byte_);
	}

	/**
	 * ByteToLong parses <code>byte</code> byte_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Long</code>
	 * @see com.github.cheejyg.conv.java.Byte
	 * @see com.github.cheejyg.conv.java.Byte#ByteToLong(byte byte_)
	 */
	public static java.lang.Long ByteToLong(byte byte_) {
		return Byte.ByteToLong(byte_);
	}

	/**
	 * ByteToInteger parses <code>byte</code> byte_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Integer</code>
	 * @see com.github.cheejyg.conv.java.Byte
	 * @see com.github.cheejyg.conv.java.Byte#ByteToInteger(byte byte_)
	 */
	public static java.lang.Integer ByteToInteger(byte byte_) {
		return Byte.ByteToInteger(byte_);
	}

	/**
	 * ByteToShort parses <code>byte</code> byte_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Short</code>
	 * @see com.github.cheejyg.conv.java.Byte
	 * @see com.github.cheejyg.conv.java.Byte#ByteToShort(byte byte_)
	 */
	public static java.lang.Short ByteToShort(byte byte_) {
		return Byte.ByteToShort(byte_);
	}

	/**
	 * ByteToByte returns <code>Byte</code> byte_.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Byte</code>
	 * @see com.github.cheejyg.conv.java.Byte
	 * @see com.github.cheejyg.conv.java.Byte#ByteToByte(byte byte_)
	 */
	public static java.lang.Byte ByteToByte(byte byte_) {
		return Byte.ByteToByte(byte_);
	}

	/**
	 * ByteToBoolean parses <code>byte</code> byte_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param byte_ The <code>byte</code> to be parsed.
	 * @return <code>Boolean</code>
	 * @see com.github.cheejyg.conv.java.Byte
	 * @see com.github.cheejyg.conv.java.Byte#ByteToBoolean(byte byte_)
	 */
	public static java.lang.Boolean ByteToBoolean(byte byte_) {
		return Byte.ByteToBoolean(byte_);
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Boolean conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * BooleanToString parses <code>boolean</code> boolean_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>String</code>
	 * @see com.github.cheejyg.conv.java.Boolean
	 * @see com.github.cheejyg.conv.java.Boolean#BooleanToString(boolean boolean_)
	 */
	public static java.lang.String BooleanToString(boolean boolean_) {
		return Boolean.BooleanToString(boolean_);
	}

	/**
	 * BooleanToLong parses <code>boolean</code> boolean_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Long</code>
	 * @see com.github.cheejyg.conv.java.Boolean
	 * @see com.github.cheejyg.conv.java.Boolean#BooleanToLong(boolean boolean_)
	 */
	public static java.lang.Long BooleanToLong(boolean boolean_) {
		return Boolean.BooleanToLong(boolean_);
	}

	/**
	 * BooleanToInteger parses <code>boolean</code> boolean_ to
	 * <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Integer</code>
	 * @see com.github.cheejyg.conv.java.Boolean
	 * @see com.github.cheejyg.conv.java.Boolean#BooleanToInteger(boolean boolean_)
	 */
	public static java.lang.Integer BooleanToInteger(boolean boolean_) {
		return Boolean.BooleanToInteger(boolean_);
	}

	/**
	 * BooleanToShort parses <code>boolean</code> boolean_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Short</code>
	 * @see com.github.cheejyg.conv.java.Boolean
	 * @see com.github.cheejyg.conv.java.Boolean#BooleanToShort(boolean boolean_)
	 */
	public static java.lang.Short BooleanToShort(boolean boolean_) {
		return Boolean.BooleanToShort(boolean_);
	}

	/**
	 * BooleanToByte parses <code>boolean</code> boolean_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be parsed.
	 * @return <code>Byte</code>
	 * @see com.github.cheejyg.conv.java.Boolean
	 * @see com.github.cheejyg.conv.java.Boolean#BooleanToByte(boolean boolean_)
	 */
	public static java.lang.Byte BooleanToByte(boolean boolean_) {
		return Boolean.BooleanToByte(boolean_);
	}

	/**
	 * BooleanToBoolean returns <code>Boolean</code> boolean_.
	 * 
	 * @author Cheejyg
	 * @param boolean_ The <code>boolean</code> to be returned.
	 * @return <code>Boolean</code>
	 * @see com.github.cheejyg.conv.java.Boolean
	 * @see com.github.cheejyg.conv.java.Boolean#BooleanToBoolean(boolean boolean_)
	 */
	public static java.lang.Boolean BooleanToBoolean(boolean boolean_) {
		return Boolean.BooleanToBoolean(boolean_);
	}
}
